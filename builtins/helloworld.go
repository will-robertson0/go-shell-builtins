package builtins

import (
    "fmt"
    "io"
)

func HelloWorld(w io.Writer, args ...string) error {
    _, err := fmt.Fprintln(w, "Hello, world!")
    return err
}
