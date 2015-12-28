package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "time"
)

// Maps strings (names) to lyprocesses.
var processes = make(map[string]*lyprocess)

// Counter of processes that are running.
// This is NOT the length of the processes map. Some processes may have exited,
// and thus are not running.
var nRunning = 0;

func main() {
	fmt.Println("Welcome to Ly!")

	reader := bufio.NewReader(os.Stdin)

	for {
        // Read and tokenize user input
        // todo: support command history
        fmt.Print("ly > ")
		line, _ := reader.ReadString('\n')
		words := strings.Fields(line)

        // No command given
        if len(words) == 0 {
            continue
        }

		switch words[0] {
		    case "new":
                // Start a new process
		        if len(words) < 3 {
		            fmt.Println("Not enough arguments")
		            break
		        }
		        startProcess(words[1], words[2:]...)
            case "kill":
                // Kill an existing process
                if len(words) < 2 {
                    fmt.Println("Not enough arguments")
                    break
                }
                killProcess(words[1])
            case "out":
                // Print the output of a process
                if len(words) < 2 {
                    fmt.Println("Not enough arguments")
                    break
                }
                printOut(words[1])
            case "in":
                // Send some input to a process
                fmt.Println("Not implemented yet")
            case "list":
                // List running processes
                list()
            case "exit":
                // Kill all processes and exit
                exit()
	        default:
                printUsage()
		}

        fmt.Println()
	}
}

func startProcess(name string, cmd ...string) {
    if processRunning(name) {
        fmt.Println("Process", name, "is already running.")
        return
    }

    cmdString := strings.Join(cmd, " ")
    fmt.Println("Starting", name, ":", cmdString)
    processes[name] = newLyprocess(cmdString)

    go func() {
        processes[name].Running = true
        nRunning++
        processes[name].Cmd.Run()
        processes[name].Running = false
        nRunning--
        // Todo: add checking for errors
    }()
}

func killProcess(name string) {
    if !processExists(name) {
        fmt.Println("Process", name, "does not exist.")
    } else {
        processes[name].Cmd.Process.Kill()
    }
}

func printOut(name string) {
    if !processExists(name) {
        fmt.Println("Process", name, "does not exist.")
    } else {
        processes[name].Stdout.Do(func(line interface{}) {
            fmt.Println(line)
        })
    }
}

func list() {
    if len(processes) == 0 {
        fmt.Println("No processes")
    } else {
        fmt.Println(len(processes), "processes:")

        for k, v := range(processes) {
            fmt.Printf("  %v(%v)", k, v.Cmd.Process.Pid)
            if !processRunning(k) {
                fmt.Println(" -- Exited")
            } else {
                fmt.Println()
            }
        }
    }
}

func exit() {
    // Try to kill all processes for 3 seconds, then exit forcefully.
    go func() {
        for k, _ := range(processes) {
            killProcess(k)
        }
    }()

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
    new <name> <command>    Spawns a new process called <name> by running <command>.
                            All processes are started in their own shell.
    kill <name>             Kills a running process by <name>
    out <name>              Outputs the most recent standard output/error for a process.
    list                    Lists running processes with their PID and status.
    exit                    Nicely quits all processes and exists.
    help                    Prints this help message.`

    fmt.Println(helpStr)
}
