package main

import (
	"fmt"
)

var memory [0x1000]uint8   // Represents the vm's 4kb of RAM
var registers [0x10]uint16 // The 16 2 bit registers

func main() {
	for _, b := range memory {
		fmt.Printf("%b", b)
	}
	fmt.Println(len(memory))
	fmt.Println(memory[0x200])
	fmt.Println(registers)
}
