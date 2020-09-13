package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BCC
func Test0x90(t *testing.T) {
	m := raw(0x90, 0xff)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)

	m = raw(0x90, 2)
	m.Tick()
	require.Equal(t, uint16(2), m.PC)

	m = raw(0x90, 2)
	m.SetCarry()
	m.Tick()
	require.Equal(t, uint16(1), m.PC)
}
