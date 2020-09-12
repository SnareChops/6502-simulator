package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// DEC a
func Test0xce(t *testing.T) {
	m, a := seedA(0xce, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0b), m.Fetch(a...))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0xce, 0x01)
	m.Tick()
	require.Equal(t, byte(0x00), m.Fetch(a...))
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0xce, 0x0)
	m.Tick()
	require.Equal(t, byte(0xff), m.Fetch(a...))
	require.True(t, m.N())
	require.False(t, m.Z())
}

// DEC a,x
func Test0xde(t *testing.T) {
	m, a := seedAX(0xde, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0b), m.Fetch(a...))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0xde, 0x01)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedAX(0xde, 0x0)
	m.Tick()
	require.Equal(t, byte(0xff), m.Fetch(a...))
	require.True(t, m.N())
	require.False(t, m.Z())
}

// DEC zp
func Test0xc6(t *testing.T) {
	m, a := seedZP(0xc6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0b), m.Fetch(a))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0xc6, 0x01)
	m.Tick()
	require.Equal(t, byte(0x00), m.Fetch(a))
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZP(0xc6, 0x0)
	m.Tick()
	require.Equal(t, byte(0xff), m.Fetch(a))
	require.True(t, m.N())
	require.False(t, m.Z())
}

// DEC zp,x
func Test0xd6(t *testing.T) {
	m, a := seedZPX(0xd6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0b), m.Fetch(a))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZPX(0xd6, 0x01)
	m.Tick()
	require.Equal(t, byte(0x00), m.Fetch(a))
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedZPX(0xd6, 0x0)
	m.Tick()
	require.Equal(t, byte(0xff), m.Fetch(a))
	require.True(t, m.N())
	require.False(t, m.Z())
}
