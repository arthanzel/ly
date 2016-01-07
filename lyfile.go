package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func runLyfile() {
    file, err := os.Open("lyfile")
    if err != nil {
        file, err = os.Open("Lyfile")
        if err != nil {
            fmt.Println("Can't open lyfile!")
        }
    }

    // Read all lines
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }

        line = strings.TrimSpace(line)
        runInstruction(line)
    }
}
