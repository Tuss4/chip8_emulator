package main

import (
    "fmt"
)

var memory[4096]uint8

func main () {
    fmt.Printf("%b\n", memory[0])
}