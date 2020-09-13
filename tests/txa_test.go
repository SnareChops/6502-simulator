package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TXA
func Test0x8a(t *testing.T) {
	m := raw(0x8a)
	m.X = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x8a)
	m.X = 0
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x8a)
	m.X = 0x82
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}
