package main

import "os/exec"

type lyprocess struct {
    File string
    Cmd *exec.Cmd
    Stdout *CircularArray
    Running bool
}

func newLyprocess(cmdString string) *lyprocess {
    ly := new(lyprocess)
    ly.File = cmdString
    ly.Cmd = exec.Command("bash", "-c", cmdString)
    ly.Cmd.Stdout = ly
    ly.Cmd.Stderr = ly
    ly.Stdout = NewCircularArray(250)
    ly.Running = false
    return ly
}

func (ly *lyprocess) Write(p []byte) (n int, err error) {
    // ly.Stdout = append(ly.Stdout, parseLines(p)...)

    for _, line := range(parseLines(p)) {
        ly.Stdout.Insert(line)
    }

    return len(p), nil
}
