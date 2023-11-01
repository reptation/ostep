package main

import (
    "log"
    "os"
)

func main() {
    fd, err := os.Create("/tmp/file")
    if err != nil {
        log.Fatal(err)
    }

    defer fd.Close()

    _, err = fd.WriteString("hello, world!\n")
    if err != nil {
        log.Fatal(err)
    }
}
