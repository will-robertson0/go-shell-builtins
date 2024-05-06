package builtins

import (
	"fmt"
	"os/exec"
)

func Source(args ...string) error {

    if len(args) < 1 {
        return fmt.Errorf("%w: source requires at least one argument", ErrInvalidArgCount)
    }

    sourceArgs := args[1:]

    // fmt.Println("path: ", args[0])

    cmd := exec.Command(args[0], sourceArgs...)
    // fmt.Println("running command: ", args[0], sourceArgs)

    err := cmd.Run()
    // fmt.Println("command finished")
    return err
}
