package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0xc8(t *testing.T) {
	m := raw(0xc8)
	m.Y = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0d), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xc8)
	m.Y = 0xff
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xc8)
	m.Y = 0x7f
	m.Tick()
	require.Equal(t, byte(0x80), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}
