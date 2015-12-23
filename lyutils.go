package main

import (
    
)

func existsProcess(name string) bool {
    _, ok := processes[name]
    return ok
}
