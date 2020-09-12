package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0xca(t *testing.T) {
	m := raw(0xca)
	m.X = 0x0c
	m.Tick()
	require.Equal(t, byte(0x0b), m.X)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0xca)
	m.X = 0x01
	m.Tick()
	require.Equal(t, byte(0x0), m.X)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xca)
	m.X = 0x0
	m.Tick()
	require.Equal(t, byte(0xff), m.X)
	require.True(t, m.N())
	require.False(t, m.Z())
}
