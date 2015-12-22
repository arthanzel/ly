package main

import (
    "fmt"
)

func removeExited() {
    fmt.Println(len(processes))
}

func existsProcess(name string) bool {
    _, ok := processes[name]
    return ok
}
