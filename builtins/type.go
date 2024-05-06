package builtins

import (
    "fmt"
    "os/exec"
)

func Type(builtinCmds [7]string, args ...string) error {
    isBuiltin := false

    // check that args[0] exists
    if len(args) == 0 {
        return error(nil)
    }

    // check if it's a builtin
    for _, b := range builtinCmds {
        if b == args[0] {
            isBuiltin = true
        }
    }
    if isBuiltin {
        _, err := fmt.Println(args[0], "is a shell builtin")
        return err
    }

    // check if it's in $PATH
    path, pathErr := exec.LookPath(args[0])
    if pathErr == nil {
        fmt.Println(args[0], "is", path)
        return pathErr
    }

    // can't be found
    fmt.Println("type: Could not find", args[0])
    return error(nil)
}
