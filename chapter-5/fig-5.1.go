package main

import (
    "fmt"
    "log"
    "os"
    "syscall"
)

func main() {
    fmt.Printf("Hello world (pid:%d)\n", os.Getpid())
    // syscall package is deprecated but sys has no equivalent to fork
    // We could maybe use syscall.SYS_FORK to get C behavior 
    rc, err := syscall.ForkExec(os.Args[0], os.Args,
        &syscall.ProcAttr{
            Env: append(os.Environ()),
            Sys: &syscall.SysProcAttr{
                Setsid: true,
            },
        })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("child process pid is %d\n", rc)
}
