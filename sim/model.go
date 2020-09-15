package sim

import "github.com/SnareChops/6502-simulator/bit"

// Model represents a 6502 processor
type Model struct {
	PC uint16
	*Memory
	*Registers
	*Flags
	opcodes Opcodes
}

// NewModel returns a new 6502 sim model
func NewModel() *Model {
	m := &Model{
		PC:        0x0,
		Memory:    NewMemory(),
		Registers: NewRegisters(),
		Flags:     NewFlags(),
	}
	m.opcodes = InitOpcodes(m)
	return m

}

// Tick executes an instruction
func (m *Model) Tick() {
	m.Exec(m.NextByte())
}

// Exec executes an instruction
func (m *Model) Exec(b byte) {
	m.opcodes[b]()
}

// NextByte increments the program counter and
// returns the byte at that address
func (m *Model) NextByte() byte {
	result := m.Fetch(AsBytes(m.PC)...)
	m.PC++
	return result
}

// NextWord increments the program counter and
// returns the next 2 bytes from those addresses
func (m *Model) NextWord() []byte {
	return []byte{m.NextByte(), m.NextByte()}
}

// AX returns a []byte memory address
// after adding the value in the X
// register
// a,x memory mode
func (m *Model) AX(b []byte) []byte {
	return AsBytes(AsUint16(b...) + uint16(m.X))
}

// AY returns a []byte memory address
// after adding the value in the Y
// register
// a,y memory mode
func (m *Model) AY(b []byte) []byte {
	return AsBytes(AsUint16(b...) + uint16(m.Y))
}

// ZP returns a []byte memory address
// after padding with zeros
// zp memory mode
func (m *Model) ZP(b byte) []byte {
	return []byte{b, 0x00}
}

// ZPX retuns a []byte memory address
// after adding the value in the X
// register
// zp,x memory mode
func (m *Model) ZPX(b byte) []byte {
	return AsBytes(uint16(b) + uint16(m.X))
}

// ZPY returns a []byte memory address
// adter adding the value in the Y
// register
// zp,y memory mode
func (m *Model) ZPY(b byte) []byte {
	return AsBytes(uint16(b) + uint16(m.Y))
}

// ZPIX returns a []byte memory address
// after applying indexing and addition
// with the X register
// (zp,x) memory mode
func (m *Model) ZPIX(b byte) []byte {
	n := uint16(b) + uint16(m.X)
	a1 := AsBytes(n)
	a2 := AsBytes(n + 1)
	return []byte{m.Fetch(a1...), m.Fetch(a2...)}
}

// ZPIY returns a []byte memory address
// after applying indexing and addition
// with the Y register
// (zp),y memory mode
func (m *Model) ZPIY(b byte) []byte {
	a1 := m.Fetch(b)
	a2 := m.Fetch(b + 1)
	return AsBytes(AsUint16(a1, a2) + uint16(m.Y))
}

// LDA performs an LDA operation
func (m *Model) LDA(val byte) {
	m.A = val
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// LDX performs an LDX operation
func (m *Model) LDX(val byte) {
	m.X = val
	m.SetN(int8(m.X) < 0)
	m.SetZ(m.X == 0)
}

// LDY performs an LDY operation
func (m *Model) LDY(val byte) {
	m.Y = val
	m.SetN(int8(m.Y) < 0)
	m.SetZ(m.Y == 0)
}

// STA performs an STA operation
func (m *Model) STA(a ...byte) {
	m.Set(m.A, a...)
}

// STX performs an STX operation
func (m *Model) STX(a ...byte) {
	m.Set(m.X, a...)
}

// STY performs an STY operation
func (m *Model) STY(a ...byte) {
	m.Set(m.Y, a...)
}

// ADC performs an ADC operation
func (m *Model) ADC(a byte) {
	initial := m.A
	m.A += a + m.C()
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
	m.SetC(uint8(m.A) < uint8(initial))
	m.SetV(Int8Overflows(int8(initial), int8(m.A)))
}

// SBC performs an SBC operation
func (m *Model) SBC(a byte) {
	initial := m.A
	m.A = m.A - a - m.borrow()
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
	m.SetC(uint8(m.A) < uint8(initial))
	m.SetV(Int8Overflows(int8(initial), int8(m.A)))
}

// INC performs an INC operation
func (m *Model) INC(r Resolver) {
	a, val := r(m)
	val++
	m.Set(val, a...)
	m.SetN(int8(val) < 0)
	m.SetZ(val == 0)
}

// INX performs an INX operation
func (m *Model) INX() {
	m.X++
	m.SetN(int8(m.X) < 0)
	m.SetZ(m.X == 0)
}

// INY performs an INY operation
func (m *Model) INY() {
	m.Y++
	m.SetN(int8(m.Y) < 0)
	m.SetZ(m.Y == 0)
}

// DEC performs a DEC operation
func (m *Model) DEC(r Resolver) {
	a, val := r(m)
	val--
	m.Set(val, a...)
	m.SetN(int8(val) < 0)
	m.SetZ(val == 0)
}

// DEX performs a DEX operation
func (m *Model) DEX() {
	m.X--
	m.SetN(int8(m.X) < 0)
	m.SetZ(m.X == 0)
}

// DEY performs a DEY operation
func (m *Model) DEY() {
	m.Y--
	m.SetN(int8(m.Y) < 0)
	m.SetZ(m.Y == 0)
}

// ASL performs an ASL operation
func (m *Model) ASL(r Resolver) {
	asl := func(val byte) byte {
		m.SetC(0b10000000&val != 0)
		val <<= 1
		m.SetN(int8(val) < 0)
		m.SetZ(byte(val) == 0)
		return val
	}
	if r != nil {
		a, val := r(m)
		m.Set(asl(val), a...)
	} else {
		m.A = asl(m.A)
	}
}

// LSR performs an LSR operation
func (m *Model) LSR(r Resolver) {
	lsr := func(val byte) byte {
		m.SetC(val&1 != 0)
		val >>= 1
		m.SetN(int8(val) < 0)
		m.SetZ(val == 0)
		return val
	}
	if r != nil {
		a, val := r(m)
		m.Set(lsr(val), a...)
	} else {
		m.A = lsr(m.A)
	}
}

// ROL performs a ROL operation
func (m *Model) ROL(r Resolver) {
	rol := func(val byte) byte {
		o := m.C()
		m.SetC(val&0b10000000 != 0)
		val <<= 1
		val |= o
		m.SetN(int8(val) < 0)
		m.SetZ(val == 0)
		return val
	}
	if r != nil {
		a, val := r(m)
		m.Set(rol(val), a...)
	} else {
		m.A = rol(m.A)
	}
}

// ROR performs a ROR operation
func (m *Model) ROR(r Resolver) {
	ror := func(val byte) byte {
		o := m.C()
		m.SetC(val&1 != 0)
		val >>= 1
		val |= o << 7
		m.SetN(int8(val) < 0)
		m.SetZ(val == 0)
		return val
	}
	if r != nil {
		a, val := r(m)
		m.Set(ror(val), a...)
	} else {
		m.A = ror(m.A)
	}
}

// AND performs an AND operation
func (m *Model) AND(r Resolver) {
	_, val := r(m)
	m.A &= val
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// ORA performs an ORA operation
func (m *Model) ORA(r Resolver) {
	_, val := r(m)
	m.A |= val
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// XOR performs an XOR operation
func (m *Model) XOR(r Resolver) {
	_, val := r(m)
	m.A ^= val
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// CMP performs a CMP operation
func (m *Model) CMP(r Resolver) {
	_, val := r(m)
	m.compare(m.A, val)
}

// CPX performs a CPX operation
func (m *Model) CPX(r Resolver) {
	_, val := r(m)
	m.compare(m.X, val)
}

// CPY performs a CPY operation
func (m *Model) CPY(r Resolver) {
	_, val := r(m)
	m.compare(m.Y, val)
}

// BIT performs a BIT operation
func (m *Model) BIT(r Resolver) {
	_, val := r(m)
	m.SetN(val&(1<<7) != 0)
	m.SetV(val&(1<<6) != 0)
	m.SetZ(m.A&val == 0)
}

// BCC performs a BCC operation
func (m *Model) BCC(b byte) {
	if m.C() == 0 {
		m.offsetPC(b)
	}
}

// BCS performs a BCS operation
func (m *Model) BCS(b byte) {
	if m.C() != 0 {
		m.offsetPC(b)
	}
}

// BNE performs a BNE operation
func (m *Model) BNE(b byte) {
	if !m.Z() {
		m.offsetPC(b)
	}
}

// BEQ performs a BEQ operation
func (m *Model) BEQ(b byte) {
	if m.Z() {
		m.offsetPC(b)
	}
}

// BPL performs a BPL operation
func (m *Model) BPL(b byte) {
	if !m.N() {
		m.offsetPC(b)
	}
}

// BMI performs a BMI operation
func (m *Model) BMI(b byte) {
	if m.N() {
		m.offsetPC(b)
	}
}

// BVC performs a BVC operation
func (m *Model) BVC(b byte) {
	if !m.V() {
		m.offsetPC(b)
	}
}

// BVS performs a BVS operation
func (m *Model) BVS(b byte) {
	if m.V() {
		m.offsetPC(b)
	}
}

// TAX performs a TAX operation
func (m *Model) TAX() {
	m.X = m.A
	m.SetN(int8(m.X) < 0)
	m.SetZ(m.X == 0)
}

// TXA performs a TXA operation
func (m *Model) TXA() {
	m.A = m.X
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// TAY performs a TAY operation
func (m *Model) TAY() {
	m.Y = m.A
	m.SetN(int8(m.Y) < 0)
	m.SetZ(m.Y == 0)
}

// TYA performs a TYA operation
func (m *Model) TYA() {
	m.A = m.Y
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// TSX performs a TSX operation
func (m *Model) TSX() {
	m.X = m.SP
	m.SetN(int8(m.X) < 0)
	m.SetZ(m.X == 0)
}

// TXS performs a TXS operation
func (m *Model) TXS() {
	m.SP = m.X
	m.SetN(int8(m.SP) < 0)
	m.SetZ(m.SP == 0)
}

// PHA performs a PHA operation
func (m *Model) PHA() {
	m.Push(m.A)
}

// PLA performs a PLA operation
func (m *Model) PLA() {
	m.A = m.Pop()
	m.SetN(int8(m.A) < 0)
	m.SetZ(m.A == 0)
}

// PHP performs a PHP operation
func (m *Model) PHP() {
	m.Push(m.SR.Byte() & 0b11001111)
}

// PLP performs a PLP operation
func (m *Model) PLP() {
	m.SR = bit.FromByte(m.SR.Byte() | m.Pop())
}

// JMP performs a JMP operation
func (m *Model) JMP(r JMPResolver) {
	val := r(m)
	m.PC = AsUint16(val...)
}

// JSR performs a JSR operation
func (m *Model) JSR(b ...byte) {
	a := AsBytes(m.PC - 1)
	m.Push(a[1])
	m.Push(a[0])
	m.PC = AsUint16(a...)
}

// RTS performs an RTS operation
func (m *Model) RTS() {
	m.PC = AsUint16(m.Pop(), m.Pop()) + 1
}

// RTI performs an RTI operation
func (m *Model) RTI() {
	m.SR = bit.FromByte(m.Pop())
	m.RTS()
}

// CLC performs a CLC operation
func (m *Model) CLC() {
	m.SetC(false)
}

// SEC performs a SEC operation
func (m *Model) SEC() {
	m.SetC(true)
}

// CLD performs a CLD operation
func (m *Model) CLD() {
	m.SetD(false)
}

// SED performs a SED operation
func (m *Model) SED() {
	m.SetD(true)
}

// CLI performs a CLI operation
func (m *Model) CLI() {
	m.SetI(false)
}

// SEI performs an SEI operation
func (m *Model) SEI() {
	m.SetI(true)
}

// CLV performs a CLV operation
func (m *Model) CLV() {
	m.SetV(false)
}

// BRK performs a BRK operation
func (m *Model) BRK() {
	m.SetB(true)
	m.SetI(true)
}

// NOP does nothing
func (m *Model) NOP() {}

func (m *Model) offsetPC(b byte) {
	b--
	if int8(b) < 0 {
		m.PC -= uint16(^b + 1)
	} else {
		m.PC += uint16(b)
	}
}

func (m *Model) compare(a, b byte) {
	m.SetN(a < b)
	m.SetZ(a == b)
	m.SetC(a >= b)
}
