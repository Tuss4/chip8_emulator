// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import (
	"fmt"
)

var memory [0x1000]uint8 // Represents the vm's 4kb of RAM
var register [0x10]uint8 // The registers
var stack [0x10]uint16   // The 16x2byte stack
var I uint16
var PC uint16
var sp uint8

func op_8xy2(op_code uint16) {
	return nil
}

func main() {
	for _, b := range memory {
		fmt.Printf("%b", b)
	}
	fmt.Println(len(memory))
	fmt.Println(memory[0x200])
	fmt.Println(stack)
	test := 0x6c3f
	fmt.Printf("%b\n", test)
	test_2 := test << 4
	fmt.Printf("%b\n", test_2)
}
