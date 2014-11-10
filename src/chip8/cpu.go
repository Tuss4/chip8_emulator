// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import "fmt"

var memory [0x1000]uint8 // Represents the vm's 4kb of RAM
var register [0x10]uint8 // The registers
var stack [0x10]uint16   // The 16x2byte stack
var I uint16
var PC uint16
var sp uint8

// To look at the highest bit >> 12
// To look at the lowest bit & 0xf

// Op_code w/ the highest bits == 0x7
func Op_7xkk(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	register[x] += kk
}

// Op_ocodes w/ the highest bits == 0x8
func Op_8xy0(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[y]
}

func Op_8xy1(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[x] | register[y]
}

func Op_8xy2(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] & register[y]
}

func Op_8xy3(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] ^ register[y]
}

func Op_8xy4(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] += register[y]
	if register[x] > 255 {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
}

func Op_8xy5(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] -= register[y]
	if register[x] > register[y] {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
}

func Op_8xy6(op_code uint16, register *[0x10]uint8) {
	x := (op_code >> 8) & 0xF
	/*
		VF = Vx & 0x1
		Vx = Vx >> 0x1
	*/
	register[x] = register[x] >> 1
	if register[x]&0x1 == 1 {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
}

func main() {
	Op_8xy2(0x8b02, &register)
	Op_8xy3(0x8d13, &register)
	fmt.Println(register)
}
