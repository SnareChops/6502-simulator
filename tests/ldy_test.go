package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// LDY a
func Test0xac(t *testing.T) {
	m, _ := seedA(0xac, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xac, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xac, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDY a,x
func Test0xbc(t *testing.T) {
	m, _ := seedAX(0xbc, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0xbc, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAX(0xbc, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDY immediate
func Test0xa0(t *testing.T) {
	m := raw(0xa0, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xa0, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xa0, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDY zp
func Test0xa4(t *testing.T) {
	m, _ := seedZP(0xa4, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xa4, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xa4, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDY zp,x
func Test0xb4(t *testing.T) {
	m, _ := seedZPX(0xb4, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0xb4, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPX(0xb4, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}
