package main

import (
	"bufio"
	"fmt"
    "io"
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
		fmt.Println("Welcome to Ly!")

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
        processes[name] = newLyprocess(cmd)

        go func() {
            bufOut := bufio.NewReader(*processes[name].Stdout)

            processes[name].Cmd.Start()
            // log.Println("Process", name, "has ended.")

            go func() {
                for {
                    line, _, err := bufOut.ReadLine()
                    if err != nil {
                        log.Println("Readline error:", err)
                        break
                    }
                    log.Println("Result", line)
                }
            }()


            // for processes[name].Cmd.ProcessStatus

            processes[name].Cmd.Wait()
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
    Stdout *io.ReadCloser
}

func newLyprocess(cmdString string) *lyprocess {
    ly := &lyprocess { cmdString, exec.Command("bash", "-c", cmdString), nil }
    stdout, _ := ly.Cmd.StdoutPipe()
    ly.Stdout = &stdout
    return ly
}
