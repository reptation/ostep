package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Need an arg")
    }

    for {
        fmt.Println(os.Args[1])
	time.Sleep(1 * time.Second)
    }
}
