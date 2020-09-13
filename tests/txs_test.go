package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TXS
func Test0x9a(t *testing.T) {
	m := raw(0x9a)
	m.X = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0c), m.SP)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x9a)
	m.X = 0
	m.Tick()
	require.Equal(t, byte(0), m.SP)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x9a)
	m.X = 0x82
	m.Tick()
	require.Equal(t, byte(0x82), m.SP)
	require.True(t, m.N())
	require.False(t, m.Z())
}
