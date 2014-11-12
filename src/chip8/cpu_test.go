package main

import "testing"

func TestOp_8xy2(t *testing.T) {
	var register [0x10]uint8
	code := uint16(0x8b02)
	Op_8xy2(code)
	if register[0xb] != 0xb&0x0 {
		t.Error("Expected 0, got ", register[0xb])
	}
}
