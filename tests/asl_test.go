package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ASL a
func Test0x0e(t *testing.T) {
	m, a := seedA(0x0e, 0b1)
	m.Tick()
	require.Equal(t, byte(0b10), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0x0e, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0x0e, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())
}

// ASL a,x
func Test0x1e(t *testing.T) {
	m, a := seedAX(0x1e, 0b1)
	m.Tick()
	require.Equal(t, byte(0b10), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0x1e, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0x1e, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())
}

// ASL
func Test0x0a(t *testing.T) {
	m := raw(0x0a)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10), m.A)
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x0a)
	m.A = 0b01000000
	m.Tick()
	require.Equal(t, byte(0b10000000), m.A)
	require.Equal(t, byte(0x0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0x0a)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())
}

// ASL zp
func Test0x06(t *testing.T) {
	m, a := seedZP(0x06, 0b1)
	m.Tick()
	require.Equal(t, byte(0b10), m.Fetch(a))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0x1e, 0b01000000)
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a))
	require.Equal(t, byte(0x0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0x1e, 0b10000000)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())
}
