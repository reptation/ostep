package main

import (
    "fmt"
    "os"
)

func main() {
    p := 32
    fmt.Printf("(%d) Memory address of p is %p\n", os.Getpid(), &p)
    // Lines excluded because Go does not have pointer arithmetic
}
