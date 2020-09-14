package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// RTS
func Test0x60(t *testing.T) {
	m := raw(0x60)
	m.Push(0xbb)
	m.Push(0xaa)
	m.Tick()
	require.Equal(t, uint16(0xbbab), m.PC)
}
