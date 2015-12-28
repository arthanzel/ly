package main

import "os/exec"
import "strings"

// Lyprocess is a structure that wraps an os/exec.Cmd object and provides some
// extra information for Ly's convenience. Ly will usually interact through a
// lyprocess rather than a Cmd.
type lyprocess struct {
    File string // The executable name
    Cmd *exec.Cmd
    Running bool

    // Circular arrays stop programs that create an endless output stream from
    // eating memory.
    Stdout *CircularArray
}

func newLyprocess(cmdString string) *lyprocess {
    ly := new(lyprocess)

    ly.File = cmdString

    // Create the Cmd object and wire its standard out and error streams to the
    // Lyprocess object.
    ly.Cmd = exec.Command("bash", "-c", cmdString)
    ly.Cmd.Stdout = ly
    ly.Cmd.Stderr = ly

    // 300 lines of output seems like a reasonable amount
    ly.Stdout = NewCircularArray(300)

    ly.Running = false

    return ly
}

// Write converts the contents of a byte buffer into a set of strings and stores
// them in the Lyprocess. This method will consume the standard output and error
// streams of running processes.
func (ly *lyprocess) Write(buffer []byte) (n int, err error) {
    // Split the byte stream into a list of lines. Remove the last blank line.
    lines := strings.Split(string(buffer), "\n")
    lines = lines[0:len(lines) - 1]

    for _, line := range(lines) {
        ly.Stdout.Insert(line)
    }

    // All of the lines were read without errors.
    return len(buffer), nil
}
