package sim_test

import (
	"testing"

	"github.com/SnareChops/6502-simulator/sim"
	"github.com/stretchr/testify/require"
)

func TestFetchSet(t *testing.T) {
	m := sim.NewMemory()
	require.Equal(t, byte(0x0), m.Fetch(0x00, 0x12))
	m.Set(0x0c, 0x00, 0x012)
	require.Equal(t, byte(0x0c), m.Fetch(0x00, 0x12))
}
