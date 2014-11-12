package main

import "fmt"
import "testing"

func TestOp_00EE(t *testing.T) {
	code := uint16(0x00EE)
	Op_00EE(code)
	if PC != stack[uint16(sp+uint8(1))] {
		t.Error("Wrong value was set to the counter.")
	}
}

func TestOp_1nnn(t *testing.T) {
	code := uint16(0x1ccc)
	low_bits := code & 0xfff
	Op_1nnn(code)
	// PC should be equal to low_bits
	if PC != low_bits {
		t.Error("Wrong value set to the counter.")
	}
}

func TestOp_2nnn(t *testing.T) {
	code := uint16(0x2bb5)
	low_bits := code & 0xfff
	Op_2nnn(code)
	// PC should be equal to low_bits
	if PC != low_bits {
		t.Error("Wrong value set to counter")
	}
}

func TestOp_8xy2(t *testing.T) {
	var register [0x10]uint8
	code := uint16(0x8b02)
	Op_8xy2(code)
	if register[0xb] != 0xb&0x0 {
		t.Error("Expected 0, got ", register[0xb])
	}
}
