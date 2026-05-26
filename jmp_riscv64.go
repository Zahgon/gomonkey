package gomonkey

// buildJmpDirective 为 riscv64 架构生成一段跳转指令，
// 将传入的 64 位地址加载到寄存器 x6（t1）中，
// 然后执行 JALR x0, 0(x6) 实现无条件跳转。
func buildJmpDirective(double uintptr) []byte {
	_ = "STUB: not implemented"

	// 将地址转换为 64 位无符号整数
	return nil
}

// 将 64 位地址拆分成若干部分：
// imm0: 位 [7:0]
// imm1: 位 [19:8] (12 位)
// imm2: 位 [31:20] (12 位)
// imm3: 位 [43:32] (12 位)
// imm4: 位 [63:44] (20 位)

// 依次生成指令：
// 1. LUI x6, imm4  // 将最高 20 位加载到 x6

// 2. ADDI x6, x6, imm3  // 加载接下来的 12 位

// 3. SLLI x6, x6, 12   // 左移 12 位

// 4. ADDI x6, x6, imm2  // 加载接下来的 12 位

// 5. SLLI x6, x6, 12   // 再次左移 12 位

// 6. ADDI x6, x6, imm1  // 加载接下来的 12 位

// 7. SLLI x6, x6, 8    // 左移 8 位

// 8. ORI x6, x6, imm0  // 最后加载最低 8 位

// 9. JALR x0, 0(x6)    // 跳转到 x6 指定的地址

// 以下辅助函数生成各条指令的机器码，均返回 4 字节小端表示。

// LUI 指令格式：
// 31          12 11   7 6       0
// [ imm[31:12] ] [ rd ] [ opcode ]
// opcode LUI 为 0x37。
func encodeLUI(rd int, imm20 uint32) []byte { _ = "STUB: not implemented"; return nil }

// ADDI 指令格式（用于加载 12 位立即数）：
// 31         20 19   15 14 12 11   7 6       0
// [ imm[11:0] ] [ rs1 ] [funct3] [ rd ] [ opcode ]
// opcode ADDI 为 0x13，funct3 为 0。
func encodeADDI(rd, rs1 int, imm int32) []byte { _ = "STUB: not implemented"; return nil }

// SLLI 指令格式：
// 31      26 25    20 19   15 14 12 11   7 6       0
// [ 0 ] [ shamt ] [ rs1 ] [funct3] [ rd ] [ opcode ]
// opcode 为 0x13，funct3 为 1。
func encodeSLLI(rd, rs1, shamt int) []byte { _ = "STUB: not implemented"; return nil }

// ORI 指令格式：
// 31         20 19   15 14 12 11   7 6       0
// [ imm[11:0] ] [ rs1 ] [funct3] [ rd ] [ opcode ]
// opcode 为 0x13，funct3 为 6。
func encodeORI(rd, rs1 int, imm int32) []byte { _ = "STUB: not implemented"; return nil }

// JALR 指令格式：
// 31         20 19   15 14 12 11   7 6       0
// [ imm[11:0] ] [ rs1 ] [funct3] [ rd ] [ opcode ]
// opcode 为 0x67，funct3 为 0。
// JALR x0, 0(x6) 用于无条件跳转。
func encodeJALR(rd, rs1 int, imm int32) []byte { _ = "STUB: not implemented"; return nil }
