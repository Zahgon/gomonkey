//go:build loong64
// +build loong64

package gomonkey

const (
	REG_R0  uint32 = 0
	REG_R29        = 29
	REG_R30        = 30
)

const (
	OP_ORI    uint32 = 0x00E << 22
	OP_LU12IW        = 0x00A << 25
	OP_LU32ID        = 0x00B << 25
	OP_LU52ID        = 0x00C << 22
	OP_LDD           = 0x0A3 << 22
	OP_JIRL          = 0x013 << 26
)

func buildJmpDirective(double uintptr) []byte { _ = "STUB: not implemented"; return nil }

// lu12i.w r29, bit31_12
// ori     r29, r29, bit11_0
// lu32i.d r29, bit51_32
// lu52i.d r29, bit63_52
// ld.d,   r30, r29, 0
// jirl    r0,  r30, 0

func wireup_opc(opc uint32, rd, rj uint32, val uintptr) []byte {
	_ = "STUB: not implemented"
	return nil
}

// rd
// rj
// si12

// rd
// si20

// rd
// rj
// si16
