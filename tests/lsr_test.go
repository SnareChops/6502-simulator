package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// LSR a
func Test0x4e(t *testing.T) {
	m, a := seedA(0x4e, 0b1)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x4e, 0b10)
	m.Tick()
	require.Equal(t, byte(0x1), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// LSR a,x
func Test0x5e(t *testing.T) {
	m, a := seedA(0x5e, 0b1)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x5e, 0b10)
	m.Tick()
	require.Equal(t, byte(0x1), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// LSR
func Test0x4a(t *testing.T) {
	m := raw(0x4a)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x4a)
	m.A = 0b10
	m.Tick()
	require.Equal(t, byte(0x1), m.A)
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// LSR zp
func Test0x46(t *testing.T) {
	m, a := seedZP(0x46, 0b1)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZP(0x46, 0b10)
	m.Tick()
	require.Equal(t, byte(0x1), m.Fetch(a))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// LSR zp,x
func Test0x56(t *testing.T) {
	m, a := seedZPX(0x56, 0b1)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZPX(0x56, 0b10)
	m.Tick()
	require.Equal(t, byte(0x1), m.Fetch(a))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
