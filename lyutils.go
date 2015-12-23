package main

import (
    "fmt"
    "os"
)

func removeExited() {
    for k, v := range(processes) {
        fmt.Println("Found ", k)
        if v.Cmd.Process != nil {
            fmt.Println(isRunning(v.Cmd.Process.Pid))
        }

        if v.Cmd.ProcessState != nil {
            fmt.Println(v.Cmd.ProcessState)
            //delete(processes, k)
        }
    }
}

func existsProcess(name string) bool {
    _, ok := processes[name]
    return ok
}

// TODO: can't get status or running process.
// Need to use a goroutime for each process and set a code on exit
func isRunning(pid int) bool {
    _, err := os.FindProcess(pid)
    return err == nil
}
