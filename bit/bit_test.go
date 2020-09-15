package bit_test

import (
	"testing"

	"github.com/SnareChops/6502-simulator/bit"
	"github.com/stretchr/testify/require"
)

func TestFromByte(t *testing.T) {
	b := byte(0b00110011)
	m := bit.FromByte(b)
	require.Equal(t, b, m.Byte())
}
