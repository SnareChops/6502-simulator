package sim

// Memory represents a 16bit addressable
// memory for the 6502 processor
type Memory struct {
	data []byte
	SP   byte
}

// NewMemory returns a new
// clean memory structure
func NewMemory() *Memory {
	return &Memory{
		data: make([]byte, 65535),
		SP:   0xfd,
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

// Push pushes an item onto the stack
func (m *Memory) Push(b byte) {
	m.SP--
	m.Set(b, m.SP, 0x01)
}

// Pop pops an item off the stack
func (m *Memory) Pop() byte {
	result := m.Fetch(m.SP, 0x01)
	m.SP++
	return result
}
