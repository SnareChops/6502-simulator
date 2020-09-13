package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BIT a
func Test0x2c(t *testing.T) {
	m, _ := seedA(0x2c, 0b11000000)
	m.A = 0b01000000
	m.Tick()
	require.True(t, m.N())
	require.True(t, m.V())
	require.False(t, m.Z())

	m, _ = seedA(0x2c, 0b10000000)
	m.Tick()
	require.True(t, m.N())
	require.False(t, m.V())
	require.True(t, m.Z())

	m, _ = seedA(0x2c, 0b01000000)
	m.Tick()
	require.False(t, m.N())
	require.True(t, m.V())
	require.True(t, m.Z())
}

// BIT immediate
func Test0x89(t *testing.T) {
	m := raw(0x89, 0b11000000)
	m.A = 0b01000000
	m.Tick()
	require.True(t, m.N())
	require.True(t, m.V())
	require.False(t, m.Z())

	m = raw(0x89, 0b10000000)
	m.Tick()
	require.True(t, m.N())
	require.False(t, m.V())
	require.True(t, m.Z())

	m = raw(0x89, 0b01000000)
	m.Tick()
	require.False(t, m.N())
	require.True(t, m.V())
	require.True(t, m.Z())
}

// BIT zp
func Test0x24(t *testing.T) {
	m, _ := seedZP(0x24, 0b11000000)
	m.A = 0b01000000
	m.Tick()
	require.True(t, m.N())
	require.True(t, m.V())
	require.False(t, m.Z())

	m, _ = seedZP(0x24, 0b10000000)
	m.Tick()
	require.True(t, m.N())
	require.False(t, m.V())
	require.True(t, m.Z())

	m, _ = seedZP(0x24, 0b01000000)
	m.Tick()
	require.False(t, m.N())
	require.True(t, m.V())
	require.True(t, m.Z())
}
