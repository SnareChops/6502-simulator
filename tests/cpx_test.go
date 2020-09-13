package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CPX a
func Test0xec(t *testing.T) {
	m, _ := seedA(0xec, 1)
	m.X = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xec, 1)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xec, 0)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CPX immediate
func Test0xe0(t *testing.T) {
	m := raw(0xe0, 1)
	m.X = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0xe0, 1)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xe0, 0)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CPX zp
func Test0xe4(t *testing.T) {
	m, _ := seedZP(0xe4, 1)
	m.X = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xe4, 1)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xe4, 0)
	m.X = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
