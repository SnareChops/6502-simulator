package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0xe8(t *testing.T) {
	m := raw(0xe8)
	m.X = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0d), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xe8)
	m.X = 0xff
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xe8)
	m.X = 0x7f
	m.Tick()
	require.Equal(t, byte(0x80), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}
