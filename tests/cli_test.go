package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CLI
func Test0x58(t *testing.T) {
	m := raw(0x58)
	m.SR = 0xff
	require.True(t, m.I())
	m.Tick()
	require.False(t, m.I())
}
