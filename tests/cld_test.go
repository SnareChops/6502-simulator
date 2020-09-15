package tests_test

import (
	"testing"

	"github.com/SnareChops/6502-simulator/bit"
	"github.com/stretchr/testify/require"
)

// CLD
func Test0xd8(t *testing.T) {
	m := raw(0xd8)
	m.SR = bit.FromByte(0xff)
	require.True(t, m.D())
	m.Tick()
	require.False(t, m.D())
}
