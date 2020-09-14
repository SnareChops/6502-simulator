package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ROR a
func Test0x6e(t *testing.T) {
	m, a := seedA(0x6e, 0b1)
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0x6e, 0b1)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x6e, 0b10)
	m.Tick()
	require.Equal(t, byte(1), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// ROR a,x
func Test0x7e(t *testing.T) {
	m, a := seedAX(0x7e, 0b1)
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0x7e, 0b1)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a...))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedAX(0x7e, 0b10)
	m.Tick()
	require.Equal(t, byte(1), m.Fetch(a...))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// ROR
func Test0x6a(t *testing.T) {
	m := raw(0x6a)
	m.A = 0b1
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0b10000000), m.A)
	require.Equal(t, byte(1), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0x6a)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x6a)
	m.A = 0b10
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// ROR zp
func Test0x66(t *testing.T) {
	m, a := seedZP(0x66, 0b1)
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0x66, 0b1)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZP(0x66, 0b10)
	m.Tick()
	require.Equal(t, byte(1), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// ROR zp,x
func Test0x76(t *testing.T) {
	m, a := seedZPX(0x76, 0b1)
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0b10000000), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedZPX(0x76, 0b1)
	m.Tick()
	require.Equal(t, byte(0), m.Fetch(a))
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZPX(0x76, 0b10)
	m.Tick()
	require.Equal(t, byte(1), m.Fetch(a))
	require.Equal(t, byte(0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
