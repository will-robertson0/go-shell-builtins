package builtins

import (
    "fmt"
    "os"
    "strings"
)

func Export(args ...string) error {
    for _, arg := range args {
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            fmt.Printf("export: invalid argument %s\n", arg)
            continue
        }
        key, value := parts[0], parts[1]
        err := os.Setenv(key, value)
        if err != nil {
            fmt.Println("export: ", err)
            return err
        }
    }

    return nil
}
