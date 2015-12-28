package main

import "os/exec"

type lyprocess struct {
    File string
    Cmd *exec.Cmd
    Stdout []string
    Running bool
}

func newLyprocess(cmdString string) *lyprocess {
    ly := &lyprocess { cmdString, exec.Command("bash", "-c", cmdString), nil, false }
    ly.Stdout = make([]string, 0)
    ly.Cmd.Stdout = ly
    ly.Cmd.Stderr = ly
    return ly
}

func (ly *lyprocess) Write(p []byte) (n int, err error) {
    ly.Stdout = append(ly.Stdout, parseLines(p)...)
    return len(p), nil
}
