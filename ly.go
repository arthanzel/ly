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
	                //pCmd.Process.Kill()
                case "list":
                    list()
		        default:
		            log.Println(line)
			}
			
			removeExited()
		}
	}
}

func startProcess(name string, cmd string, args []string) {
    if existsProcess(name) {
        fmt.Println("Process", name, "already exists.")
        return
    }
    
    fmt.Println("Starting", name, ":", cmd)
    processes[name] = &lyprocess { cmd, exec.Command(cmd) }
    processes[name].Cmd.Start()
}

func list() {
    fmt.Println(len(processes), "processes")
}

type lyprocess struct {
    File string
    //Args []string
    Cmd *exec.Cmd
}
