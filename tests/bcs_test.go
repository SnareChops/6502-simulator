package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BCS
func Test0xb0(t *testing.T) {
	m := raw(0xb0, 0xff)
	m.SEC()
	m.Tick()
	require.Equal(t, uint16(0x0), m.PC)
	m.Tick()
	require.Equal(t, uint16(0x0), m.PC)

	m = raw(0xb0, 2)
	m.SEC()
	m.Tick()
	require.Equal(t, uint16(3), m.PC)

	m = raw(0xb0, 2)
	m.Tick()
	require.Equal(t, uint16(2), m.PC)
}
