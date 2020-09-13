package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0x4e(t *testing.T) {
	m, a := seedA(0x4e, 0b1)
	m.Tick()
	require.Equal(t, byte(0x0), m.Fetch(a...))
	require.Equal(t, byte(0x1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, a = seedA(0x4e, 0b10)
	m.Tick()
	require.Equal(t, byte(0x1), m.Fetch(a...))
	require.Equal(t, byte(0x0), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
