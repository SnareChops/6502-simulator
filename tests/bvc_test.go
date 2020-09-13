package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BVC
func Test0x50(t *testing.T) {
	m := raw(0x50, 0xff)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)

	m = raw(0x50, 2)
	m.Tick()
	require.Equal(t, uint16(2), m.PC)

	m = raw(0x50, 2)
	m.SetOverflow()
	m.Tick()
	require.Equal(t, uint16(1), m.PC)
}
