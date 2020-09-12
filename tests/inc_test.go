package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// INC a
func Test0xee(t *testing.T) {
	m, a := seedA(0xee, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0d), m.Fetch(a...))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0xee, 0x7f)
	m.Tick()
	require.Equal(t, byte(0x80), m.Fetch(a...))
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedA(0xee, 0xff)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.False(t, m.N())
	require.True(t, m.Z())
}

// INC a,x
func Test0xfe(t *testing.T) {
	m, a := seedAX(0xfe, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0d), m.Fetch(a...))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0xfe, 0x7f)
	m.Tick()
	require.Equal(t, byte(0x80), m.Fetch(a...))
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedAX(0xfe, 0xff)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.False(t, m.N())
	require.True(t, m.Z())
}

func Test0xe6(t *testing.T) {
	m, a := seedZP(0xe6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0d), m.Fetch(a))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0xe6, 0x7f)
	m.Tick()
	require.Equal(t, byte(0x80), m.Fetch(a))
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedZP(0xe6, 0xff)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a))
	require.False(t, m.N())
	require.True(t, m.Z())
}

func Test0xf6(t *testing.T) {
	m, a := seedZPX(0xf6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0d), m.Fetch(a))
	require.False(t, m.N())
	require.False(t, m.Z())

	m, a = seedZPX(0xf6, 0x7f)
	m.Tick()
	require.Equal(t, byte(0x80), m.Fetch(a))
	require.True(t, m.N())
	require.False(t, m.Z())

	m, a = seedZPX(0xf6, 0xff)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a))
	require.False(t, m.N())
	require.True(t, m.Z())
}
