package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TXA
func Test0xa8(t *testing.T) {
	m := raw(0xa8)
	m.A = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0c), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xa8)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xa8)
	m.A = 0x82
	m.Tick()
	require.Equal(t, byte(0x82), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}
