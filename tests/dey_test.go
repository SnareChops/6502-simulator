package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0x88(t *testing.T) {
	m := raw(0x88)
	m.Y = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0b), m.Y)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x88)
	m.Y = 0x01
	m.Tick()
	require.Equal(t, byte(0x0), m.Y)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x88)
	m.Y = 0x0
	m.Tick()
	require.Equal(t, byte(0xff), m.Y)
	require.True(t, m.N())
	require.False(t, m.Z())
}
