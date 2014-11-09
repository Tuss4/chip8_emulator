// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import (
	"fmt"
)

var memory [0x1000]uint8 // Represents the vm's 4kb of RAM
var register [0x10]uint8
var stack [0x10]uint16 // The 16 16 bit stack

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
