package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BMI
func Test0x30(t *testing.T) {
	m := raw(0x30, 0xff)
	m.SetN(true)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)

	m = raw(0x30, 2)
	m.SetN(true)
	m.Tick()
	require.Equal(t, uint16(0x603), m.PC)

	m = raw(0x30, 2)
	m.Tick()
	require.Equal(t, uint16(0x602), m.PC)
}
