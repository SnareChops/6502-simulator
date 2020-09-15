package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BVS
func Test0x70(t *testing.T) {
	m := raw(0x70, 0xff)
	m.SetV(true)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)
	m.Tick()
	require.Equal(t, uint16(0x600), m.PC)

	m = raw(0x70, 2)
	m.SetV(true)
	m.Tick()
	require.Equal(t, uint16(0x603), m.PC)

	m = raw(0x70, 2)
	m.Tick()
	require.Equal(t, uint16(0x602), m.PC)
}
