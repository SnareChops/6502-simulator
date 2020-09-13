package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CPY a
func Test0xcc(t *testing.T) {
	m, _ := seedA(0xcc, 1)
	m.Y = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xcc, 1)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xcc, 0)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CPY immediate
func Test0xc0(t *testing.T) {
	m := raw(0xc0, 1)
	m.Y = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0xc0, 1)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xc0, 0)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CPY zp
func Test0xc4(t *testing.T) {
	m, _ := seedZP(0xc4, 1)
	m.Y = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xc4, 1)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xc4, 0)
	m.Y = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
