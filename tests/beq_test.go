package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BEQ
func Test0xf0(t *testing.T) {
	m := raw(0xf0, 0xff)
	m.SetZero()
	m.Tick()
	require.Equal(t, uint16(0), m.PC)
	m.Tick()
	require.Equal(t, uint16(0), m.PC)

	m = raw(0xf0, 2)
	m.SetZero()
	m.Tick()
	require.Equal(t, uint16(3), m.PC)

	m = raw(0xf0, 2)
	m.Tick()
	require.Equal(t, uint16(2), m.PC)
}
