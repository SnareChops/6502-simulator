package tests_test

import (
	"testing"

	"github.com/SnareChops/6502-simulator/bit"
	"github.com/stretchr/testify/require"
)

// PHP
func Test0x08(t *testing.T) {
	m := raw(0x08)
	m.SR = bit.FromByte(0xff)
	m.Tick()
	require.Equal(t, byte(0b11001111), m.Pop())
}
