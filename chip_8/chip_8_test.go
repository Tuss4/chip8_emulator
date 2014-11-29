package chip_8

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
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := register[x]
	Op_8xy2(code)
	if register[x] != old_val&register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val&register[y], ", got ", register[x])
	}
}

func TestOp_8xy3(t *testing.T) {
	code := uint16(0x8d23)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := register[x]
	Op_8xy3(code)
	if register[x] != old_val^register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val^register[y], ", got ", register[x])
	}
}

func TestOp_8xy4(t *testing.T) {
	code := uint16(0x8e14)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := register[x]
	Op_8xy4(code)
	VF := register[0xF]
	if register[x] != old_val+register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val+register[y], ", got ", register[x])
	}

	if register[x] > 255 {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_8xy5(t *testing.T) {
	code := uint16(0x8b35)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := register[x]
	Op_8xy5(code)
	VF := register[0xF]
	if register[x] != old_val-register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val-register[y], ", got ", register[x])
	}

	if register[x] > register[y] {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_8xy6(t *testing.T) {
	code := uint16(0x8126)
	x := (code >> 8) & 0xf
	old_val := register[x]
	Op_8xy6(code)
	VF := register[0xF]
	if old_val&0x1 == uint8(1) {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}

	if register[x] != old_val>>1 {
		t.Error("Incorrect value for Vx. Expected: ", old_val>>1, ", got ", register[x])
	}
}

func TestOp_8xy7(t *testing.T) {
	code := uint16(0x8a37)
	x := (code >> 8) & 0xF
	y := (code >> 4) & 0xF
	old_val := register[x]
	Op_8xy7(code)
	VF := register[0xF]
	if register[y] > register[x] {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}

	if register[x] != register[y]-old_val {
		t.Error("Incorrect value for Vx. Expected: ", register[y]-old_val, ", got ", register[x])
	}
}

func TestOp_8xyE(t *testing.T) {
	code := uint16(0x8b2E)
	x := (code >> 8) & 0xF
	old_val := register[x]
	Op_8xyE(code)
	VF := register[0xF]
	if old_val&0x8 == uint8(1) {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
	if register[x] != old_val<<1 {
		t.Error("Incorrect value for Vx. Expected: ", old_val<<1, ", got ", register[x])
	}
}

func TestOp_9xy0(t *testing.T) {
	code := uint16(0x9b20)
	x := (code >> 8) & 0xF
	y := (code >> 4) & 0xF
	old_PC := PC
	Op_9xy0(code)
	if register[x] != register[y] {
		if PC != old_PC+uint16(2) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_Annn(t *testing.T) {
	code := uint16(0xAc45)
	nnn := code & 0xFFF
	Op_Annn(code)
	if I != nnn {
		t.Error("I set incorrectly.")
	}
}

func TestOp_Bnnn(t *testing.T) {
	code := uint16(0xB54c)
	nnn := code & 0xFFF
	Op_Bnnn(code)
	if PC != nnn+uint16(register[0x0]) {
		t.Error("PC set incorrectly.")
	}
}

func TestOp_Cxkk(t *testing.T) {
	code := uint16(0xC89d)
	x := (code >> 8) & 0xF
	old_val := register[x]
	Op_Cxkk(code)
	if register[x] == old_val {
		t.Error("Vx set incorrectly.")
	}
}
