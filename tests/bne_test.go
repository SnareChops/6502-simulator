package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BNE
func Test0xd0(t *testing.T) {
	m := raw(0xd0, 0xff)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)

	m = raw(0xd0, 2)
	m.Tick()
	require.Equal(t, uint16(0x603), m.PC)

	m = raw(0xb0, 2)
	m.SetZ(true)
	m.Tick()
	require.Equal(t, uint16(0x602), m.PC)
}
