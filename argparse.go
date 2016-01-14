package main

import "strings"

// argparse converts a string into a command and an array of arguments, similar
// to how a shell parses arguments.
// argparse("uname -m -r -p") will yield ("uname", ["-m", "-r", "-p"]).
// Quoted arguments are supported. Escaped characters or interpolation are not.
func argparse(args string) (string, []string) {
    reader := strings.NewReader(args)

    currentArg := newStringBuffer()
    argv := make([]string, 0)

    // Indicates if argparse is inside a quote and should treat spaces literally
    quote := false

    // This loop reads one rune at a time and adds it to a string buffer (currentArg).
    // If the rune is a space or some other symbol that indicates the end of an argument,
    // currentArg is appended to argv and cleared.
    // The program continues to read arguments until the end of the string.
    for {
        r, _, err := reader.ReadRune()

        // Error indicates end of string
        if err != nil {
            // This check prevents blank arguments from being added.
            if currentArg.Length > 0 {
                argv = append(argv, currentArg.String())
            }
            break
        }

        if r == ' ' {
            // Keep spaces if inside a quoted string
            if quote {
                currentArg.Add(" ")
            } else if currentArg.Length > 0 {
                argv = append(argv, currentArg.String())
                currentArg = newStringBuffer()
            }
        } else if r == '"' {
            if currentArg.Length > 0 {
                argv = append(argv, currentArg.String())
                currentArg = newStringBuffer()
            }
            quote = !quote
        } else {
            currentArg.Add(string(r))
        }
    }

    return argv[0], argv[1:]
}
