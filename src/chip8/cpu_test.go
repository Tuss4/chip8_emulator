package main

import "testing"

var COUNTER_VAL_ERROR = "Wrong value set to counter."

func TestOp_00EE(t *testing.T) {
	code := uint16(0x00EE)
	old_sp := sp
	Op_00EE(code)
	if PC != stack[uint16(sp+uint8(1))] {
		t.Error(COUNTER_VAL_ERROR)
	}
	if sp != old_sp-uint8(1) {
		t.Error("Stack pointer was not properly decremented.")
	}
}

func TestOp_1nnn(t *testing.T) {
	code := uint16(0x1ccc)
	low_bits := code & 0xfff
	Op_1nnn(code)
	// PC should be equal to low_bits
	if PC != low_bits {
		t.Error(COUNTER_VAL_ERROR)
	}
}

func TestOp_2nnn(t *testing.T) {
	code := uint16(0x2bb5)
	old_sp := sp
	low_bits := code & 0xfff
	Op_2nnn(code)
	// PC should be equal to low_bits
	if PC != low_bits {
		t.Error(COUNTER_VAL_ERROR)
	}
	if sp != old_sp+uint8(1) {
		t.Error("Stack point was not properly incremented")
	}
}

func TestOp_3xkk(t *testing.T) {
	code := uint16(0x31cc)
	x_bit := (code >> 8) & 0xf
	kk := code & 0xff
	old_PC := PC
	Op_3xkk(code)
	if uint16(register[x_bit]) == kk {
		if PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_4xkk(t *testing.T) {
	code := uint16(0x4b33)
	x_bits := (code >> 8) & 0xf
	kk := code & 0xff
	old_PC := PC
	Op_4xkk(code)
	if uint16(register[x_bits]) != kk {
		if PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_5xy0(t *testing.T) {
	code := uint16(0x5b40)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_PC := PC
	Op_5xy0(code)
	if register[x] == register[y] {
		if PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_6xkk(t *testing.T) {
	code := uint16(0x6a55)
	x := (code >> 8) & 0xf
	kk := code & 0xff
	Op_6xkk(code)
	if register[x] != uint8(kk) {
		t.Error("Incorrect value for register. Expected: ", uint8(kk), "got: ", register[x])
	}
}

func TestOp_7xkk(t *testing.T) {
	code := uint16(0x7b44)
	x := (code >> 8) & 0xf
	kk := code & 0xff
	old_val := register[x]
	Op_7xkk(code)
	if register[x] != old_val+uint8(kk) {
		t.Error(
			"Incorrect value for register. Expected: ", old_val+uint8(kk), "got: ", register[x])
	}
}

func TestOp_8xy0(t *testing.T) {
	code := uint16(0x8be0)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	Op_8xy0(code)
	if register[x] != register[y] {
		t.Error("Incorrect value for Vx. Expected: ", register[y], ", got: ", register[x])
	}
}

func TestOp_8xy1(t *testing.T) {
	code := uint16(0x8b81)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := register[x]
	Op_8xy1(code)
	if register[x] != old_val|register[y] {
		t.Error(
			"Incorrect value for Vx. Expected: ", old_val|register[y], ", got: ", register[x])
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
