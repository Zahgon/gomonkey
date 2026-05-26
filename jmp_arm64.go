//go:build arm64
// +build arm64

package gomonkey

func buildJmpDirective(double uintptr) []byte { _ = "STUB: not implemented"; return nil }

// MOVZ x26, double[16:0]
// MOVK x26, double[32:16]
// MOVK x26, double[48:32]
// MOVK x26, double[64:48]
// LDR x10, [x26]
// BR x10

func movImm(opc, shift int, val uintptr) []byte { _ = "STUB: not implemented"; return nil }

// rd
// imm16
// hw
// const
// opc
// sf
