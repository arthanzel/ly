package main

// todo: add process error checking
// todo: add check for invalid command
// todo: clean up arg parsing!! (recursive?)
// todo: command history with arrow keys (not possible)
// todo: output unread lines after input (delay?)

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "time"
)

// Maps strings (names) to lyprocesses.
var processes = make(map[string]*lyprocess)

// Counter of processes that are running to keep O(1) time.
// This is NOT the length of the processes map. Some processes may have exited,
// and thus are not running.
var nRunning = 0;

func main() {
	fmt.Println("Welcome to Ly! Type 'help' if you're lost.")

	reader := bufio.NewReader(os.Stdin)

    // Run the lyfile if the -l switch is provided
    if len(os.Args) > 1 && os.Args[1] == "-l" {
        runLyfile()
    }

	for {
        // Read and tokenize user input
        fmt.Print("ly > ")
		line, _ := reader.ReadString('\n')
        runInstruction(line)

        fmt.Println()
	}
}

func runInstruction(instr string) {
    words := strings.Fields(instr)

    // No command given
    if len(words) == 0 {
        return
    }

    switch words[0] {
        case "new":
            // Start a new process
            if len(words) < 3 {
                fmt.Println("Not enough arguments.")
                break
            }
            startProcess(words[1], words[2:]...)
        case "kill":
            // Kill an existing process
            if len(words) < 2 {
                fmt.Println("Not enough arguments.")
                break
            }
            killProcess(words[1])
        case "killall":
            fmt.Println("Attempting to kill all processes.")
            killAllProcesses()
        case "out":
            // Print the output of a process
            if len(words) < 2 {
                fmt.Println("Not enough arguments.")
                break
            }
            printOut(words[1])
        case "in":
            // Send some input to a process
            if len(words) < 3 {
                fmt.Println("Not enough arguments.")
                break
            }
            sendInput(words[1], words[2:]...)
        case "list":
            // List running processes
            list()
        case "lyfile":
            runLyfile()
        case "exit":
            // Kill all processes and exit
            exit()
        default:
            printUsage()
    }
}

func startProcess(name string, cmd ...string) {
    if processRunning(name) {
        fmt.Println("Process", name, "is already running.")
        return
    }

    cmdString := strings.Join(cmd, " ")
    fmt.Printf("Starting %s :: %s.\n", name, cmdString)
    processes[name] = newLyprocess(cmdString)

    go func() {
        processes[name].Running = true
        nRunning++
        processes[name].Run()
        processes[name].Running = false
        nRunning--
    }()
}

func killAllProcesses() {
    go func() {
        for k, _ := range(processes) {
            killProcess(k)
        }
    }()
}

func killProcess(name string) {
    if !processExists(name) {
        fmt.Println("Process", name, "does not exist.")
    } else if processes[name].Cmd.Process != nil {
        processes[name].Cmd.Process.Kill()
    }
}

func printOut(name string) {
    if !processExists(name) {
        fmt.Println("Process", name, "does not exist.")
    } else {
        processes[name].PrintLog()
    }
}

func sendInput(name string, in ...string) {
    if !processExists(name) {
        fmt.Println("Process", name, "does not exist.")
    } else {
        processes[name].WriteInput(strings.Join(in, " "))
    }
}

func list() {
    if len(processes) == 0 {
        fmt.Println("No processes")
    } else {
        fmt.Println(len(processes), "processes:")

        for processName, lyproc := range(processes) {
            // Don't print the PID if the process isn't started
            if lyproc.Cmd.Process == nil {
                fmt.Printf("  %v(-)\n", processName)
            } else {
                // Name and PID notification
                fmt.Printf("  %v(%v)", processName, lyproc.Cmd.Process.Pid)

                // Unread messages notification
                if lyproc.UnreadLogLines > 1 {
                    fmt.Printf(yellowString("  %v new message"), lyproc.UnreadLogLines)
                } else if lyproc.UnreadLogLines == 1 {
                    fmt.Print(yellowString("  1 new message"))
                }

                // Exited notification
                if !processRunning(processName) {
                    fmt.Println(" -- Exited")
                } else {
                    fmt.Println()
                }
            }
        }
    }
}

func exit() {
    // Try to kill all processes for 3 seconds, then exit forcefully.
    killAllProcesses()

    // Try to exit for 3000 ms if no processes are running
    for t := 0; t < 3000; {
        if nRunning == 0 {
            fmt.Println("Goodbye!")
            os.Exit(0)
        }

        t += 300;
        time.Sleep(300 * time.Millisecond)
    }

    fmt.Println("Ly couldn't close all processes in time. These processes may still be running:")
    for k, v := range(processes) {
        if v.Running {
            fmt.Println("  " + k)
        }
    }
    fmt.Println("Goodbye!")
    os.Exit(1)
}

func printUsage() {
    helpStr := `Usage: <operation> [arguments]
Operations:
    lyfile                  Runs commands from 'Lyfile' or 'lyfile'.
    new <name> <command>    Spawns a new process called <name> by running <command>.
                            All processes are started in their own shell.
    kill <name>             Kills a running process by <name>
    killall                 Attempts to kill all running processes.
    out <name>              Outputs the most recent standard output/error for a process.
    list                    Lists running processes with their PID and status.
    exit                    Nicely quits all processes and exists.
    help                    Prints this help message.

Run ly with the -l switch to automatically run the lyfile.`

    fmt.Println(helpStr)
}
