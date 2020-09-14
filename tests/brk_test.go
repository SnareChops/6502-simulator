package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BRK
func Test0x00(t *testing.T) {
	m := raw(0x00)
	require.False(t, m.B())
	require.False(t, m.I())
	m.Tick()
	require.True(t, m.B())
	require.True(t, m.I())
}
