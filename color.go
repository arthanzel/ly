package main

// This file provides functions to print colourized strings using ANSI escape
// sequences.
// todo: Make sure this works on Windows. OSX and Linux should be good.

const (
    RED = "\x1b[31m"
    YELLOW = "\x1b[33m"
    DEFAULT = "\x1b[39m"
    GREY = "\x1b[90m"
)

func redString(s string) string {
    return RED + s + DEFAULT
}

func yellowString(s string) string {
    return YELLOW + s + DEFAULT
}

func greyString(s string) string {
    return GREY + s + DEFAULT
}
