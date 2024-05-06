package builtins

import (
    "fmt"
    "os"
)

// ok i know this one's cheap but you didn't say we couldn't
func Pwd() error {
    dir, err := os.Getwd()
    if err == nil {
        fmt.Println(dir)
    }
    return err
}
