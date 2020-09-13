package sim

// Opcodes represents a list of
// available opcodes, and their
// matching Executor
type Opcodes = map[byte]func()

// InitOpcodes intializes all know opcodes
// with the given model
func InitOpcodes(m *Model) Opcodes {
	return map[byte]func(){
		0xad: func() { m.LDA(m.Fetch(m.NextWord()...)) },
		0xbd: func() { m.LDA(m.Fetch(m.AX(m.NextWord())...)) },
		0xb9: func() { m.LDA(m.Fetch(m.AY(m.NextWord())...)) },
		0xa9: func() { m.LDA(m.NextByte()) },
		0xa5: func() { m.LDA(m.Fetch(m.ZP(m.NextByte())...)) },
		0xa1: func() { m.LDA(m.Fetch(m.ZPIX(m.NextByte())...)) },
		0xb5: func() { m.LDA(m.Fetch(m.ZPX(m.NextByte())...)) },
		0xb1: func() { m.LDA(m.Fetch(m.ZPIY(m.NextByte())...)) },
		0xae: func() { m.LDX(m.Fetch(m.NextWord()...)) },
		0xbe: func() { m.LDX(m.Fetch(m.AY(m.NextWord())...)) },
		0xa2: func() { m.LDX(m.NextByte()) },
		0xa6: func() { m.LDX(m.Fetch(m.ZP(m.NextByte())...)) },
		0xb6: func() { m.LDX(m.Fetch(m.ZPY(m.NextByte())...)) },
		0xac: func() { m.LDY(m.Fetch(m.NextWord()...)) },
		0xbc: func() { m.LDY(m.Fetch(m.AX(m.NextWord())...)) },
		0xa0: func() { m.LDY(m.NextByte()) },
		0xa4: func() { m.LDY(m.Fetch(m.ZP(m.NextByte())...)) },
		0xb4: func() { m.LDY(m.Fetch(m.ZPX(m.NextByte())...)) },
		0x8d: func() { m.STA(m.NextWord()...) },
		0x9d: func() { m.STA(m.AX(m.NextWord())...) },
		0x99: func() { m.STA(m.AY(m.NextWord())...) },
		0x85: func() { m.STA(m.ZP(m.NextByte())...) },
		0x81: func() { m.STA(m.ZPIX(m.NextByte())...) },
		0x95: func() { m.STA(m.ZPX(m.NextByte())...) },
		0x91: func() { m.STA(m.ZPIY(m.NextByte())...) },
		0x8e: func() { m.STX(m.NextWord()...) },
		0x86: func() { m.STX(m.ZP(m.NextByte())...) },
		0x96: func() { m.STX(m.ZPY(m.NextByte())...) },
		0x8c: func() { m.STY(m.NextWord()...) },
		0x84: func() { m.STY(m.ZP(m.NextByte())...) },
		0x94: func() { m.STY(m.ZPX(m.NextByte())...) },
		0x6d: func() { m.ADC(m.Fetch(m.NextWord()...)) },
		0x7d: func() { m.ADC(m.Fetch(m.AX(m.NextWord())...)) },
		0x79: func() { m.ADC(m.Fetch(m.AY(m.NextWord())...)) },
		0x69: func() { m.ADC(m.NextByte()) },
		0x65: func() { m.ADC(m.Fetch(m.ZP(m.NextByte())...)) },
		0x61: func() { m.ADC(m.Fetch(m.ZPIX(m.NextByte())...)) },
		0x75: func() { m.ADC(m.Fetch(m.ZPX(m.NextByte())...)) },
		0x71: func() { m.ADC(m.Fetch(m.ZPIY(m.NextByte())...)) },
		0xed: func() { m.SBC(m.Fetch(m.NextWord()...)) },
		0xfd: func() { m.SBC(m.Fetch(m.AX(m.NextWord())...)) },
		0xf9: func() { m.SBC(m.Fetch(m.AY(m.NextWord())...)) },
		0xe9: func() { m.SBC(m.NextByte()) },
		0xe5: func() { m.SBC(m.Fetch(m.ZP(m.NextByte())...)) },
		0xe1: func() { m.SBC(m.Fetch(m.ZPIX(m.NextByte())...)) },
		0xf5: func() { m.SBC(m.Fetch(m.ZPX(m.NextByte())...)) },
		0xf1: func() { m.SBC(m.Fetch(m.ZPIY(m.NextByte())...)) },
		0xee: func() { m.INC(AResolver(m.NextWord()...)) },
		0xfe: func() { m.INC(AXResolver(m.NextWord()...)) },
		0xe6: func() { m.INC(ZPResolver(m.NextByte())) },
		0xf6: func() { m.INC(ZPXResolver(m.NextByte())) },
		0xe8: func() { m.INX() },
		0xc8: func() { m.INY() },
		0xce: func() { m.DEC(AResolver(m.NextWord()...)) },
		0xde: func() { m.DEC(AXResolver(m.NextWord()...)) },
		0xc6: func() { m.DEC(ZPResolver(m.NextByte())) },
		0xd6: func() { m.DEC(ZPXResolver(m.NextByte())) },
		0xca: func() { m.DEX() },
		0x88: func() { m.DEY() },
		0x0e: func() { m.ASL(AResolver(m.NextWord()...)) },
		0x1e: func() { m.ASL(AXResolver(m.NextWord()...)) },
		0x0a: func() { m.ASL(nil) },
		0x06: func() { m.ASL(ZPResolver(m.NextByte())) },
		0x16: func() { m.ASL(ZPXResolver(m.NextByte())) },
		0x4e: func() { m.LSR(AResolver(m.NextWord()...)) },
		0x5e: func() { m.LSR(AXResolver(m.NextWord()...)) },
		0x4a: func() { m.LSR(nil) },
		0x46: func() { m.LSR(ZPResolver(m.NextByte())) },
		0x56: func() { m.LSR(ZPXResolver(m.NextByte())) },
		0x2e: func() { m.ROL(AResolver(m.NextWord()...)) },
		0x3e: func() { m.ROL(AXResolver(m.NextWord()...)) },
		0x2a: func() { m.ROL(nil) },
		0x26: func() { m.ROL(ZPResolver(m.NextByte())) },
		0x36: func() { m.ROL(ZPXResolver(m.NextByte())) },
		0x6e: func() { m.ROR(AResolver(m.NextWord()...)) },
		0x7e: func() { m.ROR(AXResolver(m.NextWord()...)) },
		0x6a: func() { m.ROR(nil) },
		0x66: func() { m.ROR(ZPResolver(m.NextByte())) },
		0x76: func() { m.ROR(ZPXResolver(m.NextByte())) },
		0x2d: func() { m.AND(AResolver(m.NextWord()...)) },
		0x3d: func() { m.AND(AXResolver(m.NextWord()...)) },
		0x39: func() { m.AND(AYResolver(m.NextWord()...)) },
		0x29: func() { m.AND(IResolver(m.NextByte())) },
		0x25: func() { m.AND(ZPResolver(m.NextByte())) },
		0x21: func() { m.AND(ZPIXResolver(m.NextByte())) },
		0x35: func() { m.AND(ZPXResolver(m.NextByte())) },
		0x31: func() { m.AND(ZPIYResolver(m.NextByte())) },
		0x0d: func() { m.ORA(AResolver(m.NextWord()...)) },
		0x1d: func() { m.ORA(AXResolver(m.NextWord()...)) },
		0x19: func() { m.ORA(AYResolver(m.NextWord()...)) },
		0x09: func() { m.ORA(IResolver(m.NextByte())) },
		0x05: func() { m.ORA(ZPResolver(m.NextByte())) },
		0x01: func() { m.ORA(ZPIXResolver(m.NextByte())) },
		0x15: func() { m.ORA(ZPXResolver(m.NextByte())) },
		0x11: func() { m.ORA(ZPIYResolver(m.NextByte())) },
		0x4d: func() { m.XOR(AResolver(m.NextWord()...)) },
		0x5d: func() { m.XOR(AXResolver(m.NextWord()...)) },
		0x59: func() { m.XOR(AYResolver(m.NextWord()...)) },
		0x49: func() { m.XOR(IResolver(m.NextByte())) },
		0x45: func() { m.XOR(ZPResolver(m.NextByte())) },
		0x41: func() { m.XOR(ZPIXResolver(m.NextByte())) },
		0x55: func() { m.XOR(ZPXResolver(m.NextByte())) },
		0x51: func() { m.XOR(ZPIYResolver(m.NextByte())) },
		0xcd: func() { m.CMP(AResolver(m.NextWord()...)) },
		0xdd: func() { m.CMP(AXResolver(m.NextWord()...)) },
		0xd9: func() { m.CMP(AYResolver(m.NextWord()...)) },
		0xc9: func() { m.CMP(IResolver(m.NextByte())) },
		0xc5: func() { m.CMP(ZPResolver(m.NextByte())) },
		0xc1: func() { m.CMP(ZPIXResolver(m.NextByte())) },
		0xd5: func() { m.CMP(ZPXResolver(m.NextByte())) },
		0xd1: func() { m.CMP(ZPIYResolver(m.NextByte())) },
		0xec: func() { m.CPX(AResolver(m.NextWord()...)) },
		0xe0: func() { m.CPX(IResolver(m.NextByte())) },
		0xe4: func() { m.CPX(ZPResolver(m.NextByte())) },
	}
}
