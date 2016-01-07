package main

import "strings"

func argparse(args string) (cmd string, argv []string) {
    reader := strings.NewReader(args)

    currentTerm := newStringBuffer()
    quote := false

    for {
        r, _, err := reader.ReadRune()

        // Error (usually EOF) indicates end of string
        if err != nil {
            if currentTerm.Length > 0 {
                argv = append(argv, currentTerm.String())
            }
            break
        }

        if r == ' ' {
            // Keep spaces if inside a quoted string
            if quote {
                currentTerm.Add(" ")
            } else if currentTerm.Length > 0 {
                argv = append(argv, currentTerm.String())
                currentTerm = newStringBuffer()
            }
        } else if r == '"' {
            if currentTerm.Length > 0 {
                argv = append(argv, currentTerm.String())
                currentTerm = newStringBuffer()
            }
            quote = !quote
        } else {
            currentTerm.Add(string(r))
        }
    }

    return argv[0], argv[1:]
}
