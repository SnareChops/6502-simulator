package sim

const (
	n = 7
	v = 6
	b = 4
	d = 3
	i = 2
	z = 1
	c = 0
)

// Model represents a 6502 processor
type Model struct {
	A  byte
	X  byte
	Y  byte
	PC uint16
	SP byte
	SR byte
	*Memory
	opcodes Opcodes
}

// NewModel returns a new 6502 sim model
func NewModel() *Model {
	m := &Model{
		A:      0x0,
		X:      0x0,
		Y:      0x0,
		PC:     0xffff,
		SP:     0x0,
		SR:     0b00100000,
		Memory: NewMemory(),
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
	m.PC++
	return m.Fetch(AsBytes(m.PC)...)
}

// NextWord increments the program counter and
// returns the next 2 bytes from those addresses
func (m *Model) NextWord() []byte {
	return []byte{m.NextByte(), m.NextByte()}
}

// N returns the value of the Negative register
func (m *Model) N() bool {
	return m.getRegisterBit(n) != 0
}

// V returns the value of the Overflow register
func (m *Model) V() bool {
	return m.getRegisterBit(v) != 0
}

// B returns the value of the Break register
func (m *Model) B() bool {
	return m.getRegisterBit(b) != 0
}

// D returns the value of the Decimal register
func (m *Model) D() bool {
	return m.getRegisterBit(d) != 0
}

// I returns the value of the Interupt Disable register
func (m *Model) I() bool {
	return m.getRegisterBit(i) != 0
}

// Z returns the value of the Zero register
func (m *Model) Z() bool {
	return m.getRegisterBit(z) != 0
}

// C returns the value of the Carry register
func (m *Model) C() byte {
	return m.getRegisterBit(c)
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
	if int8(val) < 0 {
		m.setRegisterBit(n)
	} else {
		m.clearRegisterBit(n)
	}
	if val == 0 {
		m.setRegisterBit(z)
	} else {
		m.clearRegisterBit(z)
	}
}

// LDX performs an LDX operation
func (m *Model) LDX(val byte) {
	m.X = val
	if int8(val) < 0 {
		m.setRegisterBit(n)
	} else {
		m.clearRegisterBit(n)
	}
	if val == 0 {
		m.setRegisterBit(z)
	} else {
		m.clearRegisterBit(z)
	}
}

// LDY performs an LDY operation
func (m *Model) LDY(val byte) {
	m.Y = val
	if int8(val) < 0 {
		m.setRegisterBit(n)
	} else {
		m.clearRegisterBit(n)
	}
	if val == 0 {
		m.setRegisterBit(z)
	} else {
		m.clearRegisterBit(z)
	}
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
	m.updateRegisterBit(n, int8(m.A) < 0)
	m.updateRegisterBit(z, m.A == 0)
	m.updateRegisterBit(c, uint8(m.A) < uint8(initial))
	m.updateRegisterBit(v, Int8Overflows(int8(initial), int8(m.A)))
}

// SBC performs an SBC operation
func (m *Model) SBC(a byte) {
	initial := m.A
	m.A = m.A - a - m.borrow()
	m.updateRegisterBit(n, int8(m.A) < 0)
	m.updateRegisterBit(z, m.A == 0)
	m.updateRegisterBit(c, uint8(m.A) < uint8(initial))
	m.updateRegisterBit(v, Int8Overflows(int8(initial), int8(m.A)))
}

// INC performs an INC operation
func (m *Model) INC(r Resolver) {
	a, val := r(m)
	val++
	m.Set(val, a...)
	m.updateRegisterBit(n, int8(val) < 0)
	m.updateRegisterBit(z, val == 0)
}

// INX performs an INX operation
func (m *Model) INX() {
	m.X++
	m.updateRegisterBit(n, int8(m.X) < 0)
	m.updateRegisterBit(z, m.X == 0)
}

// INY performs an INY operation
func (m *Model) INY() {
	m.Y++
	m.updateRegisterBit(n, int8(m.Y) < 0)
	m.updateRegisterBit(z, m.Y == 0)
}

// DEC performs a DEC operation
func (m *Model) DEC(r Resolver) {
	a, val := r(m)
	val--
	m.Set(val, a...)
	m.updateRegisterBit(n, int8(val) < 0)
	m.updateRegisterBit(z, val == 0)
}

// DEX performs a DEX operation
func (m *Model) DEX() {
	m.X--
	m.updateRegisterBit(n, int8(m.X) < 0)
	m.updateRegisterBit(z, m.X == 0)
}

// DEY performs a DEY operation
func (m *Model) DEY() {
	m.Y--
	m.updateRegisterBit(n, int8(m.Y) < 0)
	m.updateRegisterBit(z, m.Y == 0)
}

// ASL performs an ASL operation
func (m *Model) ASL(r Resolver) {
	var result uint16
	if r != nil {
		a, val := r(m)
		result = uint16(val) << 1
		m.Set(byte(result), a...)
	} else {
		result = uint16(m.A) << 1
		m.A = byte(result)
	}
	m.updateRegisterBit(c, 0b100000000&result != 0)
	m.updateRegisterBit(n, int8(result) < 0)
	m.updateRegisterBit(z, byte(result) == 0)
}

// SetCarry sets the carry flag
func (m *Model) SetCarry() {
	m.setRegisterBit(c)
}

// ClearCarry clears the carry flag
func (m *Model) ClearCarry() {
	m.clearRegisterBit(c)
}

func (m *Model) borrow() byte {
	if m.C() == 0x0 {
		return 0x01
	}
	return 0x0
}

func (m *Model) getRegisterBit(n int8) byte {
	return ((1 << n) & m.SR)
}

func (m *Model) updateRegisterBit(n int8, v bool) {
	if v {
		m.setRegisterBit(n)
	} else {
		m.clearRegisterBit(n)
	}
}

func (m *Model) setRegisterBit(n int8) {
	m.SR |= 1 << n
}

func (m *Model) clearRegisterBit(n int8) {
	m.SR &^= 1 << n
}
