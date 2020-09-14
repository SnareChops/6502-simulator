package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// PLA
func Test0x68(t *testing.T) {
	m := raw(0x68)
	m.Push(0x88)
	m.Tick()
	require.Equal(t, byte(0x88), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0x68)
	m.Push(0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x68)
	m.Push(0x0)
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}
