package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ROL a
func Test0x2e(t *testing.T) {
	m, a := seedA(0x2e, 0b1)
	m.SetCarry()
	m.Tick()
	require.Equal(t, byte(0b11), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0x2e, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x2e, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ROL a,x
func Test0x3e(t *testing.T) {
	m, a := seedA(0x3e, 0b1)
	m.SetCarry()
	m.Tick()
	require.Equal(t, byte(0b11), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0x3e, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x3e, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ROL
func Test0x2a(t *testing.T) {
	m := raw(0x2a)
	m.A = 0b1
	m.SetCarry()
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x2a)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x2a)
	m.A = 0b01000000
	m.Tick()
	require.Equal(t, byte(0b10000000), m.A)
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ROL zp
func Test0x26(t *testing.T) {
	m, a := seedZP(0x26, 0b1)
	m.SetCarry()
	m.Tick()
	require.Equal(t, byte(0b11), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0x26, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZP(0x26, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ROL zp,x
func Test0x36(t *testing.T) {
	m, a := seedZPX(0x36, 0b1)
	m.SetCarry()
	m.Tick()
	require.Equal(t, byte(0b11), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZPX(0x36, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZPX(0x36, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())
}
