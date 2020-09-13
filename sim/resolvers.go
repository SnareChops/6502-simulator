package sim

// Resolver a function that resolves a memory
// address
type Resolver = func(m *Model) ([]byte, byte)

// AResolver returns a resolver that
// resolves an absolute memory address
func AResolver(b ...byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		return b, m.Fetch(b...)
	}
}

// AXResolver returns a resolver that
// resolves an a,x memory address
func AXResolver(b ...byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		address := AsBytes(AsUint16(b...) + uint16(m.X))
		return address, m.Fetch(address...)
	}
}

// AYResolver returns a resolver that
// resolves an a,y memory address
func AYResolver(b ...byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		address := AsBytes(AsUint16(b...) + uint16(m.Y))
		return address, m.Fetch(address...)
	}
}

// ZPResolver returns a resolver that
// resolves a zero page memory address
func ZPResolver(b byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		address := []byte{b, 0x00}
		return address, m.Fetch(address...)
	}
}

// ZPXResolver returns a resolver that
// resolves a zp,x memory address
func ZPXResolver(b byte) Resolver {
	return AXResolver(b)
}

// ZPIXResolver returns a resolver that
// resolves a (zp,x) memory address
func ZPIXResolver(b byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		i := uint16(b) + uint16(m.X)
		a1 := m.Fetch(AsBytes(i)...)
		a2 := m.Fetch(AsBytes(i + 1)...)
		return []byte{a1, a2}, m.Fetch(a1, a2)
	}
}

// ZPIYResolver returns a resolver that
// resolves a (zp),y memory address
func ZPIYResolver(b byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		a1 := m.Fetch(b)
		a2 := m.Fetch(b + 1)
		a := AsBytes(AsUint16(a1, a2) + uint16(m.Y))
		return a, m.Fetch(a...)
	}
}

// IResolver returns a resolver that
// returns an immediate value
func IResolver(b byte) Resolver {
	return func(m *Model) ([]byte, byte) {
		return nil, b
	}
}
