package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var processes = make(map[string]*lyprocess)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Start REPL
		log.Println("Welcome to Ly!")

		reader := bufio.NewReader(os.Stdin)

		for {
            fmt.Print("ly > ")
			line, _ := reader.ReadString('\n')
			words := strings.Fields(line)

			switch words[0] {
    			case "quit":
    			    return
    		    case "new":
    		        if len(words) < 3 {
    		            fmt.Println("Not enough arguments")
    		            break
    		        }
    		        startProcess(words[1], words[2], words[2:])
	            case "kill":
	                if len(words) < 2 {
                        fmt.Println("Not enough arguments")
                        break
                    }
                    killProcess(words[1])
                case "list":
                    list()
		        default:
		            log.Println(line)
			}

			// removeExited()
            fmt.Println()
		}
	}
}

func startProcess(name string, cmd string, args []string) {
    if existsProcess(name) {
        fmt.Println("Process", name, "already exists.")
    } else {
        log.Println("Starting", name, ":", cmd)
        processes[name] = &lyprocess { cmd, exec.Command(cmd) }

        go func() {
            processes[name].Cmd.Run() // Blocking
            log.Println("Process", name, "has ended.")
            delete(processes, name)
            // Todo: add checking for errors
        }()
    }
}

func killProcess(name string) {
    if !existsProcess(name) {
        fmt.Println("Process", name, "does not exist.")
    } else {
        log.Println("Killing", name)
        processes[name].Cmd.Process.Kill()
    }
}

func list() {
    if len(processes) == 0 {
        fmt.Println("No processes")
    } else {
        fmt.Println(len(processes), "processes:")

        for k, v := range(processes) {
            fmt.Printf("  %v (%v)\n", k, v.Cmd.Process.Pid)
        }
    }
}

type lyprocess struct {
    File string
    //Args []string
    Cmd *exec.Cmd
}
