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
