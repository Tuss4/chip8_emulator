package main

import (
    "fmt"
)

var memory [4096]uint8 // Represents the vm's 4kb of RAM

func main () {
    for _, b := range memory {
        fmt.Printf("%b", b)
    }
    fmt.Println(len(memory))
}
