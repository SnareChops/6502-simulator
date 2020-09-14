package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BPL
func Test0x10(t *testing.T) {
	m := raw(0x10, 0xff)
	m.Tick()
	require.Equal(t, uint16(0), m.PC)
	m.Tick()
	require.Equal(t, uint16(0), m.PC)

	m = raw(0x10, 2)
	m.Tick()
	require.Equal(t, uint16(3), m.PC)

	m = raw(0x10, 2)
	m.SetNegative()
	m.Tick()
	require.Equal(t, uint16(2), m.PC)
}
