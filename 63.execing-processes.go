package main

import (
    "os/exec"
    "os"
    "syscall"
)

func main() {
    binary, err := exec.LookPath("ls")
    if err != nil {
        panic(err)
    }

    args := []string{"ls", "-a", "-h", "-l"}

    env := os.Environ()

    err = syscall.Exec(binary, args, env)
    if err != nil {
        panic(err)
    }
}
