package chip_8

import "testing"

var COUNTER_VAL_ERROR = "Wrong value set to counter."
var t_c CPU

func TestOp_00EE(t *testing.T) {
	code := uint16(0x00EE)
	old_sp := t_c.sp
	t_c.Op_00EE(code)
	if t_c.PC != t_c.stack[uint16(t_c.sp+uint8(1))] {
		t.Error(COUNTER_VAL_ERROR)
	}
	if t_c.sp != old_sp-uint8(1) {
		t.Error("Stack pointer was not properly decremented.")
	}
}

func TestOp_1nnn(t *testing.T) {
	code := uint16(0x1ccc)
	low_bits := code & 0xfff
	t_c.Op_1nnn(code)
	// PC should be equal to low_bits
	if t_c.PC != low_bits {
		t.Error(COUNTER_VAL_ERROR)
	}
}

func TestOp_2nnn(t *testing.T) {
	code := uint16(0x2bb5)
	old_sp := t_c.sp
	low_bits := code & 0xfff
	t_c.Op_2nnn(code)
	// PC should be equal to low_bits
	if t_c.PC != low_bits {
		t.Error(COUNTER_VAL_ERROR)
	}
	if t_c.sp != old_sp+uint8(1) {
		t.Error("Stack point was not properly incremented")
	}
}

func TestOp_3xkk(t *testing.T) {
	code := uint16(0x31cc)
	x_bit := (code >> 8) & 0xf
	kk := code & 0xff
	old_PC := t_c.PC
	t_c.Op_3xkk(code)
	if uint16(t_c.register[x_bit]) == kk {
		if t_c.PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_4xkk(t *testing.T) {
	code := uint16(0x4b33)
	x_bits := (code >> 8) & 0xf
	kk := code & 0xff
	old_PC := t_c.PC
	t_c.Op_4xkk(code)
	if uint16(t_c.register[x_bits]) != kk {
		if t_c.PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_5xy0(t *testing.T) {
	code := uint16(0x5b40)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_PC := t_c.PC
	t_c.Op_5xy0(code)
	if t_c.register[x] == t_c.register[y] {
		if t_c.PC != old_PC+uint16(2) {
			t.Error(COUNTER_VAL_ERROR)
		}
	}
}

func TestOp_6xkk(t *testing.T) {
	code := uint16(0x6a55)
	x := (code >> 8) & 0xf
	kk := code & 0xff
	t_c.Op_6xkk(code)
	if t_c.register[x] != uint8(kk) {
		t.Error("Incorrect value for register. Expected: ", uint8(kk), "got: ", t_c.register[x])
	}
}

func TestOp_7xkk(t *testing.T) {
	code := uint16(0x7b44)
	x := (code >> 8) & 0xf
	kk := code & 0xff
	old_val := t_c.register[x]
	t_c.Op_7xkk(code)
	if t_c.register[x] != old_val+uint8(kk) {
		t.Error(
			"Incorrect value for register. Expected: ", old_val+uint8(kk), "got: ", t_c.register[x])
	}
}

func TestOp_8xy0(t *testing.T) {
	code := uint16(0x8be0)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	t_c.Op_8xy0(code)
	if t_c.register[x] != t_c.register[y] {
		t.Error("Incorrect value for Vx. Expected: ", t_c.register[y], ", got: ", t_c.register[x])
	}
}

func TestOp_8xy1(t *testing.T) {
	code := uint16(0x8b81)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy1(code)
	if t_c.register[x] != old_val|t_c.register[y] {
		t.Error(
			"Incorrect value for Vx. Expected: ", old_val|t_c.register[y], ", got: ",
			t_c.register[x])
	}
}

func TestOp_8xy2(t *testing.T) {
	code := uint16(0x8b02)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy2(code)
	if t_c.register[x] != old_val&t_c.register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val&t_c.register[y], ", got ",
			t_c.register[x])
	}
}

func TestOp_8xy3(t *testing.T) {
	code := uint16(0x8d23)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy3(code)
	if t_c.register[x] != old_val^t_c.register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val^t_c.register[y], ", got ",
			t_c.register[x])
	}
}

func TestOp_8xy4(t *testing.T) {
	code := uint16(0x8e14)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy4(code)
	VF := t_c.register[0xF]
	if t_c.register[x] != old_val+t_c.register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val+t_c.register[y], ", got ",
			t_c.register[x])
	}

	if t_c.register[x] > 255 {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_8xy5(t *testing.T) {
	code := uint16(0x8b35)
	x := (code >> 8) & 0xf
	y := (code >> 4) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy5(code)
	VF := t_c.register[0xF]
	if t_c.register[x] != old_val-t_c.register[y] {
		t.Error("Incorrect value for Vx. Expected: ", old_val-t_c.register[y], ", got ",
			t_c.register[x])
	}

	if t_c.register[x] > t_c.register[y] {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_8xy6(t *testing.T) {
	code := uint16(0x8126)
	x := (code >> 8) & 0xf
	old_val := t_c.register[x]
	t_c.Op_8xy6(code)
	VF := t_c.register[0xF]
	if old_val&0x1 == uint8(1) {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}

	if t_c.register[x] != old_val>>1 {
		t.Error("Incorrect value for Vx. Expected: ", old_val>>1, ", got ", t_c.register[x])
	}
}

func TestOp_8xy7(t *testing.T) {
	code := uint16(0x8a37)
	x := (code >> 8) & 0xF
	y := (code >> 4) & 0xF
	old_val := t_c.register[x]
	t_c.Op_8xy7(code)
	VF := t_c.register[0xF]
	if t_c.register[y] > t_c.register[x] {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}

	if t_c.register[x] != t_c.register[y]-old_val {
		t.Error("Incorrect value for Vx. Expected: ", t_c.register[y]-old_val, ", got ",
			t_c.register[x])
	}
}

func TestOp_8xyE(t *testing.T) {
	code := uint16(0x8b2E)
	x := (code >> 8) & 0xF
	old_val := t_c.register[x]
	t_c.Op_8xyE(code)
	VF := t_c.register[0xF]
	if old_val&0x8 == uint8(1) {
		if VF != uint8(1) {
			t.Error("Flag not set.")
		}
	}
	if t_c.register[x] != old_val<<1 {
		t.Error("Incorrect value for Vx. Expected: ", old_val<<1, ", got ", t_c.register[x])
	}
}

func TestOp_9xy0(t *testing.T) {
	code := uint16(0x9b20)
	x := (code >> 8) & 0xF
	y := (code >> 4) & 0xF
	old_PC := t_c.PC
	t_c.Op_9xy0(code)
	if t_c.register[x] != t_c.register[y] {
		if t_c.PC != old_PC+uint16(2) {
			t.Error("Flag not set.")
		}
	}
}

func TestOp_Annn(t *testing.T) {
	code := uint16(0xAc45)
	nnn := code & 0xFFF
	t_c.Op_Annn(code)
	if t_c.I != nnn {
		t.Error("I set incorrectly.")
	}
}

func TestOp_Bnnn(t *testing.T) {
	code := uint16(0xB54c)
	nnn := code & 0xFFF
	t_c.Op_Bnnn(code)
	if t_c.PC != nnn+uint16(t_c.register[0x0]) {
		t.Error("PC set incorrectly.")
	}
}

func TestOp_Cxkk(t *testing.T) {
	code := uint16(0xC89d)
	x := (code >> 8) & 0xF
	old_val := t_c.register[x]
	t_c.Op_Cxkk(code)
	if t_c.register[x] == old_val {
		t.Error("Vx set incorrectly.")
	}
}
