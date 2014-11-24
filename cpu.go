// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

var memory [0x1000]uint8 // Represents the vm's 4kb of RAM
var register [0x10]uint8 // The registers
var stack [0x10]uint16   // The 16x2byte stack
var I uint16
var PC uint16
var sp uint8
var DT uint8 // Delay Timer
var ST uint8 // Sound Timer
var rom_path string

// To look at the highest bits >> 12
// To look at the lowest bits & 0xf

// Op_codes w/ the highest bits == 0x0
// func Op_00E0() {
// 	// TODO: set up gfx then clear them with this function
// 	return nil
// }

func RunCPU() {
	fmt.Println(PC, sp, stack, register, I)
	for {
		code := (uint16(memory[PC]) << 8) | uint16(memory[PC+uint16(1)])
		switch {
		case code == 0x00EE:
			Op_00EE(code)
		case (code >> 12) == 0x1:
			Op_1nnn(code)
		case (code >> 12) == 0x2:
			Op_2nnn(code)
		case (code >> 12) == 0x3:
			Op_3xkk(code)
		case (code >> 12) == 0x4:
			Op_4xkk(code)
		case (code >> 12) == 0x5:
			Op_5xy0(code)
		case (code >> 12) == 0x6:
			Op_6xkk(code)
		case (code >> 12) == 0x7:
			Op_7xkk(code)
		case (code >> 12) == 0x8:
			switch {
			case (code & 0xf) == 0x0:
				Op_8xy0(code)
			case (code & 0xf) == 0x1:
				Op_8xy1(code)
			case (code & 0xf) == 0x2:
				Op_8xy2(code)
			case (code & 0xf) == 0x3:
				Op_8xy3(code)
			case (code & 0xf) == 0x4:
				Op_8xy4(code)
			case (code & 0xf) == 0x5:
				Op_8xy5(code)
			case (code & 0xf) == 0x6:
				Op_8xy6(code)
			case (code & 0xf) == 0x7:
				Op_8xy7(code)
			case (code & 0xf) == 0xE:
				Op_8xyE(code)
			}
		case (code >> 12) == 0x9:
			Op_9xy0(code)
		case (code >> 12) == 0xA:
			Op_Annn(code)
		case (code >> 12) == 0xB:
			Op_Bnnn(code)
		case (code >> 12) == 0xC:
			Op_Cxkk(code)
		case (code >> 12) == 0xD:
			Op_Dxyn(code)
		default:
			fmt.Printf("Opcode: %#x not implemented.\n", code)
		}
	}
	fmt.Println(PC, sp, stack, register, I)
}

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
	PC += uint16(2)
}

// Op_code w/ the highest bits == 0x7
func Op_7xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	register[x] += uint8(kk)
	PC += uint16(2)
}

// Op_ocodes w/ the highest bits == 0x8
func Op_8xy0(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[y]
	PC += uint16(2)
}

func Op_8xy1(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	register[x] = register[x] | register[y]
	PC += uint16(2)
}

func Op_8xy2(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] & register[y]
	PC += uint16(2)
}

func Op_8xy3(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] = register[x] ^ register[y]
	PC += uint16(2)
}

func Op_8xy4(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] += register[y]
	if register[x] > 255 {
		register[0xF] = uint8(1)
	} else {
		register[0xF] = uint8(0)
	}
	PC += uint16(2)
}

func Op_8xy5(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	register[x] -= register[y]
	if register[x] > register[y] {
		register[0xF] = uint8(1)
	} else {
		register[0xF] = uint8(0)
	}
	PC += uint16(2)
}

func Op_8xy6(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if register[x]&0x1 == uint8(1) {
		register[0xF] = uint8(1)
	} else {
		register[0xF] = uint8(0)
	}
	register[x] = register[x] >> 1
	PC += uint16(2)
}

func Op_8xy7(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	if register[y] > register[x] {
		register[0xF] = uint8(1)
	} else {
		register[0xF] = uint8(0)
	}
	register[x] = register[y] - register[x]
	PC += uint16(2)
}

func Op_8xyE(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if register[x]&0x8 == uint8(1) {
		register[0xF] = uint8(1)
	} else {
		register[0xF] = uint8(0)
	}
	register[x] = register[x] << 1
	PC += uint16(2)
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
func Op_Annn(op_code uint16) {
	nnn := op_code & 0xfff
	I = nnn
	PC += uint16(2)
}

func Op_Bnnn(op_code uint16) {
	nnn := op_code & 0xfff
	PC = nnn + uint16(register[0x0])
}

func Op_Cxkk(op_code uint16) {
	x := (op_code >> 8) & 0xF
	kk := op_code & 0xff
	rand := uint16(rand.Intn(0xff))
	register[x] = uint8(rand & kk)
	PC += uint16(2)
}

func Op_Dxyn(op_code uint16) {
	PC += uint16(2)
}

func Op_Ex9E(op_code uint16) {}

func Op_ExA1(op_code uint16) {}

func Op_Fx07(op_code uint16) {}

func Op_Fx0A(op_code uint16) {}

func Op_Fx15(op_code uint16) {}

func Op_Fx18(op_code uint16) {}

func Op_Fx1E(op_code uint16) {}

func Op_Fx29(op_code uint16) {}

func Op_Fx33(op_code uint16) {}

func Op_Fx55(op_code uint16) {}

func Op_Fx65(op_code uint16) {}

func LoadGame(rom []byte) {
	for b := 0; b < len(rom); b++ {
		memory[b+0x200] = rom[b]
	}
}

func main() {
	PC = uint16(0x200)
	if len(os.Args) < 2 {
		fmt.Println("Please specify the path to a rom.")
	} else {
		rom_path = os.Args[len(os.Args)-1]
	}
	// Set up the reading of the bytes
	if rom_path != "" {
		bytes, err := ioutil.ReadFile(rom_path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Now running ", rom_path)
		LoadGame(bytes)
		RunCPU()
	} else {
		fmt.Println("No rom specified, dude.")
	}
}
