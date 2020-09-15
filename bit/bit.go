package bit

// Mask represents a byte of bits
type Mask byte

// Set sets a bit in the byte
// func Set(b byte, n uint8) byte {
func (m *Mask) Set(n uint8) {
	*m |= 1 << n
	// return b
}

// Clear clears the bit in the byte
// func Clear(b byte, n uint8) byte {
func (m *Mask) Clear(n uint8) {
	*m &^= 1 << n
}

// Get gets the value of a bit in the byte
func (m *Mask) Get(n uint8) bool {
	return 1<<n&*m != 0
}

// Update updates the value of a bit in the byte
// func Update(b byte, n uint8, v bool) byte {
func (m *Mask) Update(n uint8, v bool) {
	if v {
		m.Set(n)
	} else {
		m.Clear(n)
	}
}

// Byte returns the mask as a byte
func (m *Mask) Byte() byte {
	return byte(*m)
}

// FromByte creates a new Mask from a byte
func FromByte(b byte) *Mask {
	m := new(Mask)
	for i := uint8(0); i < 8; i++ {
		m.Update(i, b&(1<<i) != 0)
	}
	return m
}
