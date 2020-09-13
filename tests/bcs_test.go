package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BCS
func Test0xb0(t *testing.T) {
	m := raw(0xb0, 0xff)
	m.SetCarry()
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)

	m = raw(0xb0, 2)
	m.SetCarry()
	m.Tick()
	require.Equal(t, uint16(2), m.PC)

	m = raw(0xb0, 2)
	m.Tick()
	require.Equal(t, uint16(1), m.PC)
}
