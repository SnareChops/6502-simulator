package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TAX
func Test0xaa(t *testing.T) {
	m := raw(0xaa)
	m.A = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0c), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xaa)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xaa)
	m.A = 0x82
	m.Tick()
	require.Equal(t, byte(0x82), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}
