package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// STX a
func Test0x8e(t *testing.T) {
	m, a := seedA(0x8e, 0x0)
	m.X = 0x10
	require.Equal(t, byte(0x0), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x10), m.Fetch(a...))
}

// STX zp
func Test0x86(t *testing.T) {
	m, a := seedZP(0x86, 0x0)
	m.X = 0x10
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x10), m.Fetch(a))
}

// STX zp,y
func Test0x96(t *testing.T) {
	m, a := seedZPY(0x96, 0x0)
	m.X = 0x10
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x10), m.Fetch(a))
}
