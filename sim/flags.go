package sim

import "github.com/SnareChops/6502-simulator/bit"

// Flags represents the status register flags
// of the 6502 processor
type Flags struct {
	SR *bit.Mask
}

// NewFlags returns a new Flags object
func NewFlags() *Flags {
	m := new(bit.Mask)
	m.Set(5)
	return &Flags{m}
}

// N returns the value of the Negative register
func (f *Flags) N() bool {
	return f.SR.Get(7)
}

// V returns the value of the Overflow register
func (f *Flags) V() bool {
	return f.SR.Get(6)
}

// B returns the value of the Break register
func (f *Flags) B() bool {
	return f.SR.Get(4)
}

// D returns the value of the Decimal register
func (f *Flags) D() bool {
	return f.SR.Get(3)
}

// I returns the value of the Interupt Disable register
func (f *Flags) I() bool {
	return f.SR.Get(2)
}

// Z returns the value of the Zero register
func (f *Flags) Z() bool {
	return f.SR.Get(1)
}

// C returns the value of the Carry register
func (f *Flags) C() byte {
	if f.SR.Get(0) {
		return 1
	}
	return 0
}

// SetN sets the N register bit
// to the provided bit value
func (f *Flags) SetN(val bool) {
	f.SR.Update(7, val)
}

// SetV sets the V register bit
// to the provided bit value
func (f *Flags) SetV(val bool) {
	f.SR.Update(6, val)
}

// SetB sets the B register bit
// to the provided bit value
func (f *Flags) SetB(val bool) {
	f.SR.Update(4, val)
}

// SetD sets the D register bit
// to the provided bit value
func (f *Flags) SetD(val bool) {
	f.SR.Update(3, val)
}

// SetI sets the I register bit
// to the provided bit value
func (f *Flags) SetI(val bool) {
	f.SR.Update(2, val)
}

// SetZ sets the Z register bit
// to the provided bit value
func (f *Flags) SetZ(val bool) {
	f.SR.Update(1, val)
}

// SetC sets the C register bit
// to the provided bit value
func (f *Flags) SetC(val bool) {
	f.SR.Update(0, val)
}

func (f *Flags) borrow() byte {
	if f.C() == 0x0 {
		return 0x1
	}
	return 0x0
}
