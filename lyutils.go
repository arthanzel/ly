package main

import (
    "strings"
)

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

func parseLines(buffer []byte) []string {
    source := strings.Split(string(buffer), "\n")
    lines := make([]string, 0)

    // todo: this code trims blank lines that are supposed to be there
    for _, l := range(source) {
        trimmed := strings.TrimSpace(l)
        if len(trimmed) > 0 {
            lines = append(lines, l)
        }
    }

    return lines
}
