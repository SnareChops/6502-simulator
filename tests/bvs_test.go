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
	require.Equal(t, uint16(0), m.PC)
	m.Tick()
	require.Equal(t, uint16(0), m.PC)

	m = raw(0x70, 2)
	m.SetV(true)
	m.Tick()
	require.Equal(t, uint16(3), m.PC)

	m = raw(0x70, 2)
	m.Tick()
	require.Equal(t, uint16(2), m.PC)
}
