package sim

// Stack represents a program stack
type Stack struct {
	*Memory
	SP byte
}

// NewStack returns a new stack for
// the provided memory
func NewStack(m *Memory) *Stack {
	return &Stack{SP: 0xfd, Memory: m}
}

// Push pushes an item onto the stack
func (s *Stack) Push(b byte) {
	s.SP--
	s.Set(b, s.SP, 0x01)
}

// Pop pops an item off the stack
func (s *Stack) Pop() byte {
	result := s.Fetch(s.SP, 0x01)
	s.SP++
	return result
}
