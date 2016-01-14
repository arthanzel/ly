package main

import (
    "bufio"
    "fmt"
    "io"
    "os/exec"
    "time"
)

const LOG_LENGTH = 300

// Lyprocess is a structure that wraps an os/exec.Cmd object and provides some
// extra information for Ly's convenience. Ly will usually interact through a
// lyprocess rather than a Cmd.
type lyprocess struct {
    File string // The executable name
    Argv []string
    Cmd *exec.Cmd
    Running bool
    Input io.WriteCloser

    // Circular arrays stop programs that create an endless output stream from
    // eating memory.
    Log *CircularStringArray
    UnreadLogLines int
}

func newLyprocess(cmdString string) *lyprocess {
    ly := new(lyprocess)

    ly.File, ly.Argv = argparse(cmdString)

    // Create the Cmd object and wire its standard out and error streams to the
    // Lyprocess object.
    ly.Cmd = exec.Command(ly.File, ly.Argv...)

    // 300 lines of output seems like a reasonable amount
    ly.Log = NewCircularStringArray(LOG_LENGTH)
    ly.UnreadLogLines = 0

    ly.Running = false

    return ly
}

func (ly *lyprocess) Run() {
    // Wire the stdout/error to buffers.
    // The buffers will concurrently read from the stream and add output and
    // error lines to the process's log.
    stdoutFile, stdoutErr := ly.Cmd.StdoutPipe()
    stderrFile, stderrErr := ly.Cmd.StderrPipe()
    stdin,      stdinErr  := ly.Cmd.StdinPipe()
    if stdoutErr != nil || stderrErr != nil || stdinErr != nil {
        fmt.Println("Couldn't start the process. Standard streams are misbehaving.")
        return
    }

    stdoutBuffer := bufio.NewReader(stdoutFile)
    go func() {
        for {
            line, _, err := stdoutBuffer.ReadLine()
            if err != nil {
                // EOF or closed pipe.
                // Error indicates that the process has exited.
                break
            }
            ly.WriteLine(string(line))
        }
    }()

    stderrBuffer := bufio.NewReader(stderrFile)
    go func() {
        for {
            line, _, err := stderrBuffer.ReadLine()
            if err != nil {
                break
            }
            ly.WriteErrorLine(string(line))
        }
    }()

    // Set up the standard input
    ly.Input = stdin

    ly.Cmd.Run()
    ly.Input.Write([]byte("\nfoo\n"))
}

// PrintLog prints out the process's log of outputs and errors.
func (ly *lyprocess) PrintLog() {
    ly.Log.Do(func(i int, line string) {
        if i < ly.Log.Length - ly.UnreadLogLines {
            fmt.Println(greyString(line))
        } else {
            fmt.Println(line)
        }
    })
    ly.UnreadLogLines = 0
}

func (ly *lyprocess) WriteInput(s string) {
    ly.Input.Write([]byte(s + "\n"))
}

// WriteLine adds a timestamped line of text to the process's output/error log.
func (ly *lyprocess) WriteLine(line string) {
    timeString := time.Now().Format("15:04:05.000 :: ")
    ly.Log.Insert(timeString + line)
    ly.UnreadLogLines = intmin(ly.UnreadLogLines + 1, LOG_LENGTH)
}

// WriteErrorLine adds a timestamped line of text to the process's output/error
// log and marks it with the colour red to indicate an error.
func (ly *lyprocess) WriteErrorLine(line string) {
    timeString := time.Now().Format("15:04:05.000 :: ")
    ly.Log.Insert(timeString + redString(line))
    ly.UnreadLogLines = intmin(ly.UnreadLogLines + 1, LOG_LENGTH)
}
