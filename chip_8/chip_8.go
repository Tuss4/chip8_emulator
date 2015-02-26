package chip_8

// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

import (
	// "fmt"
	"math/rand"
)

type Signal struct {
	Msg    string
	Xcoord uint8
	Ycoord uint8
	Bytes  []uint8
}

type CPU struct {
	memory   [0x1000]uint8 // Represents the vm's 4kb of RAM
	register [0x10]uint8   // The registers
	stack    [0x10]uint16  // The 16x2byte stack
	I        uint16
	PC       uint16
	sp       uint8
	DT       uint8          // Delay Timer
	ST       uint8          // Sound Timer
	gfx      [64 * 32]uint8 // Helps to keep track of pixels on the screen
}

// To look at the highest bits >> 12
// To look at the lowest bits & 0xf

var sprites = []uint8{
	0xf0, 0x90, 0x90, 0x90, 0xf0,
	0x20, 0x60, 0x20, 0x20, 0x70,
	0xf0, 0x10, 0xf0, 0x80, 0xf0,
	0xf0, 0x10, 0xf0, 0x10, 0xf0,
	0x90, 0x90, 0xf0, 0x10, 0x10,
	0xf0, 0x80, 0xf0, 0x10, 0xf0,
	0xf0, 0x80, 0xf0, 0x90, 0xf0,
	0xf0, 0x10, 0x20, 0x40, 0x40,
	0xf0, 0x90, 0xf0, 0x90, 0xf0,
	0xf0, 0x90, 0xf0, 0x10, 0xf0,
	0xf0, 0x90, 0xf0, 0x90, 0x90,
	0xe0, 0x90, 0xe0, 0x90, 0xe0,
	0xf0, 0x80, 0x80, 0x80, 0xf0,
	0xe0, 0x90, 0x90, 0x90, 0xe0,
	0xf0, 0x80, 0xf0, 0x80, 0xf0,
	0xf0, 0x80, 0xf0, 0x80, 0x80,
}

func (c *CPU) RunCPU(sig chan Signal) {
	for {
		code := (uint16(c.memory[c.PC]) << 8) | uint16(c.memory[c.PC+uint16(1)])
		switch {
		case code == 0x00E0:
			c.Op_00E0(code, sig)
		case code == 0x00EE:
			c.Op_00EE(code)
		case (code >> 12) == 0x1:
			c.Op_1nnn(code)
		case (code >> 12) == 0x2:
			c.Op_2nnn(code)
		case (code >> 12) == 0x3:
			c.Op_3xkk(code)
		case (code >> 12) == 0x4:
			c.Op_4xkk(code)
		case (code >> 12) == 0x5:
			c.Op_5xy0(code)
		case (code >> 12) == 0x6:
			c.Op_6xkk(code)
		case (code >> 12) == 0x7:
			c.Op_7xkk(code)
		case (code >> 12) == 0x8:
			switch {
			case (code & 0xf) == 0x0:
				c.Op_8xy0(code)
			case (code & 0xf) == 0x1:
				c.Op_8xy1(code)
			case (code & 0xf) == 0x2:
				c.Op_8xy2(code)
			case (code & 0xf) == 0x3:
				c.Op_8xy3(code)
			case (code & 0xf) == 0x4:
				c.Op_8xy4(code)
			case (code & 0xf) == 0x5:
				c.Op_8xy5(code)
			case (code & 0xf) == 0x6:
				c.Op_8xy6(code)
			case (code & 0xf) == 0x7:
				c.Op_8xy7(code)
			case (code & 0xf) == 0xE:
				c.Op_8xyE(code)
			}
		case (code >> 12) == 0x9:
			c.Op_9xy0(code)
		case (code >> 12) == 0xA:
			c.Op_Annn(code)
		case (code >> 12) == 0xB:
			c.Op_Bnnn(code)
		case (code >> 12) == 0xC:
			c.Op_Cxkk(code)
		case (code >> 12) == 0xD:
			c.Op_Dxyn(code, sig)
		case (code >> 12) == 0xE:
			switch {
			case (code & 0x00FF) == 0x9E:
				c.Op_Ex9E(code)
			case (code & 0x00FF) == 0xA1:
				c.Op_ExA1(code)
			}
		case (code >> 12) == 0xF:
			switch {
			case (code & 0x00FF) == 0x07:
				c.Op_Fx07(code)
			case (code & 0x00FF) == 0x0A:
				c.Op_Fx0A(code)
			case (code & 0x00FF) == 0x15:
				c.Op_Fx15(code)
			case (code & 0x00FF) == 0x18:
				c.Op_Fx18(code)
			case (code & 0x00FF) == 0x1E:
				c.Op_Fx1E(code)
			case (code & 0x00FF) == 0x29:
				c.Op_Fx29(code)
			case (code & 0x00FF) == 0x33:
				c.Op_Fx33(code)
			}
		}
	}
}

// Op_codes w/ the highest bits == 0x0
func (c *CPU) Op_00E0(op_code uint16, sig chan Signal) {
	// TODO: set up gfx then clear them with this function
	msg := Signal{"clear", 0, 0, []uint8{}}
	sig <- msg
	c.PC += uint16(2)
}

func (c *CPU) Op_00EE(op_code uint16) {
	c.PC = c.stack[uint16(c.sp)]
	c.sp -= uint8(1)
	c.PC += uint16(2)
}

// Op_code w/ the highest bits == 0x1
func (c *CPU) Op_1nnn(op_code uint16) {
	nnn := op_code & 0xfff
	c.PC = nnn
	c.PC += uint16(2)
}

// Op_code w/ the highest bits == 0x2
func (c *CPU) Op_2nnn(op_code uint16) {
	nnn := op_code & 0xfff
	c.sp += uint8(1)
	c.stack[uint16(c.sp)] = c.PC
	c.PC = nnn
	c.PC += uint16(2)
}

// Op_code w/ the highest bits == 0x3
func (c *CPU) Op_3xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	if uint16(c.register[x]) == kk {
		c.PC += uint16(2)
	} else {
		c.PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x4
func (c *CPU) Op_4xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	if uint16(c.register[x]) != kk {
		c.PC += uint16(2)
	} else {
		c.PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x5
func (c *CPU) Op_5xy0(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	if c.register[x] == c.register[y] {
		c.PC += uint16(2)
	} else {
		c.PC += uint16(2)
	}
}

// Op_code w/ the highest bits == 0x6
func (c *CPU) Op_6xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	c.register[x] = uint8(kk)
	c.PC += uint16(2)
}

// Op_code w/ the highest bits == 0x7
func (c *CPU) Op_7xkk(op_code uint16) {
	x := (op_code >> 8) & 0xf
	kk := op_code & 0xff
	c.register[x] += uint8(kk)
	c.PC += uint16(2)
}

// Op_ocodes w/ the highest bits == 0x8
func (c *CPU) Op_8xy0(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	c.register[x] = c.register[y]
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy1(op_code uint16) {
	x := (op_code >> 8) & 0xf
	y := (op_code >> 4) & 0xf
	c.register[x] = c.register[x] | c.register[y]
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy2(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	c.register[x] = c.register[x] & c.register[y]
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy3(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	c.register[x] = c.register[x] ^ c.register[y]
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy4(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	c.register[x] += c.register[y]
	if c.register[x] > 255 {
		c.register[0xF] = uint8(1)
	} else {
		c.register[0xF] = uint8(0)
	}
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy5(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	c.register[x] -= c.register[y]
	if c.register[x] > c.register[y] {
		c.register[0xF] = uint8(1)
	} else {
		c.register[0xF] = uint8(0)
	}
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy6(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if c.register[x]&0x1 == uint8(1) {
		c.register[0xF] = uint8(1)
	} else {
		c.register[0xF] = uint8(0)
	}
	c.register[x] = c.register[x] >> 1
	c.PC += uint16(2)
}

func (c *CPU) Op_8xy7(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	if c.register[y] > c.register[x] {
		c.register[0xF] = uint8(1)
	} else {
		c.register[0xF] = uint8(0)
	}
	c.register[x] = c.register[y] - c.register[x]
	c.PC += uint16(2)
}

func (c *CPU) Op_8xyE(op_code uint16) {
	x := (op_code >> 8) & 0xF
	if c.register[x]&0x8 == uint8(1) {
		c.register[0xF] = uint8(1)
	} else {
		c.register[0xF] = uint8(0)
	}
	c.register[x] = c.register[x] << 1
	c.PC += uint16(2)
}

// Op_code w/ the highest bits == 0x9
func (c *CPU) Op_9xy0(op_code uint16) {
	x := (op_code >> 8) & 0xF
	y := (op_code >> 4) & 0xF
	if c.register[x] != c.register[y] {
		c.PC += uint16(2)
	} else {
		c.PC += uint16(2)
	}
}

// OP_code w/ the highest bits == Hex letter
// Can ignore these opcodes for now.
func (c *CPU) Op_Annn(op_code uint16) {
	nnn := op_code & 0xfff
	c.I = nnn
	c.PC += uint16(2)
}

func (c *CPU) Op_Bnnn(op_code uint16) {
	nnn := op_code & 0xfff
	c.PC = nnn + uint16(c.register[0x0])
}

func (c *CPU) Op_Cxkk(op_code uint16) {
	x := (op_code >> 8) & 0xF
	kk := op_code & 0xff
	rand := uint16(rand.Intn(0xff))
	c.register[x] = uint8(rand & kk)
	c.PC += uint16(2)
}

func (c *CPU) Op_Dxyn(op_code uint16, sig chan Signal) {
	pixel := uint8(0)
	x := c.register[(op_code>>8)&0xF]
	y := c.register[(op_code>>4)&0xF]
	height := op_code & 0x000F
	msg := Signal{}
	c.register[0xF] = uint8(0)
	var gfx = make([]uint8, height)

	for yline := uint8(0); yline < uint8(height); yline++ {
		pixel = c.memory[c.I+uint16(yline)]
		for xline := uint8(0); xline < 8; xline++ {
			if pixel&(0x80>>xline) != 0 {
				if c.gfx[x+xline+((y+yline)*64)] == 1 {
					c.register[0xF] = 1
				}
				c.gfx[x+xline+((y+yline)*64)] ^= 1
			}
		}
		gfx[yline] = pixel
	}
	msg.Msg = "draw"
	msg.Xcoord = x
	msg.Ycoord = y
	msg.Bytes = gfx
	sig <- msg
	c.PC += uint16(2)
}

func (c *CPU) Op_Ex9E(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_ExA1(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx07(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx0A(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx15(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx18(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx1E(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx29(op_code uint16) {
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx33(op_code uint16) {
	x := c.register[(op_code&0x0f00)>>8]
	c.memory[c.I] = x / 100
	c.memory[c.I+uint16(1)] = (x / 10) % 10
	c.memory[c.I+uint16(2)] = (x % 100) % 10
	c.PC += uint16(2)
}

func (c *CPU) Op_Fx55(op_code uint16) {}

func (c *CPU) Op_Fx65(op_code uint16) {}

func (c *CPU) LoadGame(rom []byte) {
	for b := 0; b < len(rom); b++ {
		c.memory[b+0x200] = rom[b]
	}

	// load sprites into memory.
	for s := 0; s < len(sprites); s++ {
		c.memory[s+0x000] = sprites[s]
	}
}
