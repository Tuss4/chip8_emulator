// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import (
	"fmt"
	// "math/rand"
	"io/ioutil"
	"log"
)

var memory [0x1000]uint8 // Represents the vm's 4kb of RAM
var register [0x10]uint8 // The registers
var stack [0x10]uint16   // The 16x2byte stack
var I uint16
var PC uint16
var sp uint8

// To look at the highest bit >> 12
// To look at the lowest bit & 0xf

// Op_codes w/ the highest bits == 0x0
// func Op_00E0() {
// 	// TODO: set up gfx then clear them with this function
// 	return nil
// }

func Op_00EE(op_code uint16) {
	PC = stack[uint16(sp)]
	sp -= uint8(1)
}

// Op_code w/ the highest bits == 0x1
func Op_1nnn(op_code uint16) {
	nnn := op_code & 0xfff
	PC = nnn
}

// Op_code w/ the highest bits == 0x2
func Op_2nnn(op_code uint16) {
	nnn := op_code & 0xfff
	sp += uint8(1)
	stack[sp] = PC
	PC = nnn
}

// Op_code w/ the highest bits == 0x3
func Op_3xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	if uint16(register[x]) == kk {
		PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x4
func Op_4xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	if uint16(register[x]) != kk {
		PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x5
func Op_5xy0(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	if register[x] == register[y] {
		PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x6
func Op_6xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	register[x] = uint8(kk)
}

// Op_code w/ the highest bits == 0x7
func Op_7xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	register[x] += uint8(kk)
}

// Op_ocodes w/ the highest bits == 0x8
func Op_8xy0(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[y]
}

func Op_8xy1(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[x] | register[y]
}

func Op_8xy2(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] & register[y]
}

func Op_8xy3(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] ^ register[y]
}

func Op_8xy4(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] += register[y]
	if register[x] > 255 {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
}

func Op_8xy5(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] -= register[y]
	if register[x] > register[y] {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
}

func Op_8xy6(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if register[x]&0x1 == 1 {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[x] = register[x] >> 1
}

func Op_8xy7(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	if register[y] > register[x] {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[x] = register[y] - register[x]
}

func Op_8xyE(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if register[x]&0x8 == 1 {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[x] = register[x] << 1
}

// Op_code w/ the highest bits == 0x9
func Op_9xy0(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	if register[x] != register[y] {
		PC += uint16(2)
	}
}

// OP_code w/ the highest bits == Hex letter
// Can ignore these opcodes for now.
// func Op_Annn(op_code uint16) {
// 	nnn := op_code & 0xfff
// 	I = nnn
// }

// func Op_Bnnn(op_code uint16) {
// 	nnn := op_code & 0xfff
// 	PC = nnn + register[0x0]
// }

// func Op_Cxkk(op_code uint16) {
// 	x := (op_code >> 8) & 0xF
// 	kk := op_code & 0xff
// 	rand := uint8(rand.Intn(0xff))
// 	register[x] = rand & kk
// }

// func Op_Dxyn(op_code uint16) {
// 	return nil
// }

func main() {
	Op_8xy2(0x8b02)
	Op_8xy3(0x8d13)
	fmt.Println(register)
	// Set up the reading of the bytes
	bytes, err := ioutil.ReadFile("PONG")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes)
}
