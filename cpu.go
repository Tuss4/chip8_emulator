package main

import (
    "fmt"
)

var memory[4096]uint8 // Represents the vm's 4kb of RAM

func main () {
    fmt.Printf("%b\n", memory[0])
}