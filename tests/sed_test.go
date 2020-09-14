package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// SED
func Test0xf8(t *testing.T) {
	m := raw(0xf8)
	require.False(t, m.D())
	m.Tick()
	require.True(t, m.D())
}
