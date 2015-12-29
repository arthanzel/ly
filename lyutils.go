package main

// Returns true if a process by a given name exists. It may have exited, in
// which case its output is still accessible.
func processExists(name string) bool {
    _, ok := processes[name]
    return ok
}

func processRunning(name string) bool {
    if processExists(name) {
        return processes[name].Running
    } else {
        return false
    }
}

func intmin(a int, b int) int {
    if a <= b {
        return a
    }
    return b
}
