package main

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "time"
)

var counter = 0

func worker(loops int) {
    for i := 0; i < loops; i++ {
        counter = counter + 1
    }
}

func main() {
    if len(os.Args) != 2 {
        log.Fatal("usage: threads <value>\n")
    }

    loops, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal("Could not parse argument as integer")
    }
    go worker(loops)
    go worker(loops)
    // Add a sleep to allow Goroutines to execute
    // In "real life" we'd use channels but this is for comparison to Pthread_create
    // Interestingly, this gives the right answer everytime while the naive C equivalent does not
    time.Sleep(1 * time.Second)
    fmt.Printf("Final value: %d\n", counter)
}
