package sim

// Memory represents a 16bit addressable
// memory for the 6502 processor
type Memory struct {
	data []byte
}

// NewMemory returns a new
// clean memory structure
func NewMemory() *Memory {
	return &Memory{
		data: make([]byte, 65535),
	}
}

// Fetch retrieves a value at address b
func (m *Memory) Fetch(b ...byte) byte {
	return m.data[AsUint16(b...)]
}

// Set sets a value at address b
func (m *Memory) Set(val byte, b ...byte) {
	m.data[AsUint16(b...)] = val
}
