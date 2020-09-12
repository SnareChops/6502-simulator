package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// LDX a
func Test0xae(t *testing.T) {
	m, _ := seedA(0xae, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xae, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xae, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDX a,y
func Test0xbe(t *testing.T) {
	m, _ := seedAY(0xbe, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0xbe, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAY(0xbe, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDX immediate
func Test0xa2(t *testing.T) {
	m := raw(0xa2, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xa2, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xa2, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDX zp
func Test0xa6(t *testing.T) {
	m, _ := seedZP(0xa6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xa6, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xa6, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDX zp,y
func Test0xb6(t *testing.T) {
	m, _ := seedZPY(0xb6, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPY(0xb6, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPY(0xb6, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}
