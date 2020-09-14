package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// SEI
func Test0x78(t *testing.T) {
	m := raw(0x78)
	require.False(t, m.I())
	m.Tick()
	require.True(t, m.I())
}
