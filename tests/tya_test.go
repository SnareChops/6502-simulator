package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TYA
func Test0x98(t *testing.T) {
	m := raw(0x98)
	m.Y = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x98)
	m.Y = 0
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x98)
	m.Y = 0x82
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}
