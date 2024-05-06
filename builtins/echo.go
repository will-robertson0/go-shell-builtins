package builtins

import (
    "fmt"
    "io"
)

func Echo(w io.Writer, args ...string) error {
    err := error(nil)
    for _, c := range args {
        _, err = fmt.Fprint(w, c, " ")
    }
    println()
    return err
}
