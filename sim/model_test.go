package sim_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/SnareChops/6502-simulator/sim"
)

func TestNewModel(t *testing.T) {
	m := sim.NewModel()
	require.Equal(t, uint8(0), m.A)
	require.Equal(t, uint8(0), m.X)
	require.Equal(t, uint8(0), m.Y)
	require.Equal(t, uint16(0xffff), m.PC)
	require.Equal(t, uint8(0), m.SP)
	require.Equal(t, uint8(0b00100000), m.SR)
}

func TestN(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.N())
	m.SR = 0b10000000
	require.True(t, m.N())
}

func TestV(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.V())
	m.SR = 0b01000000
	require.True(t, m.V())
}

func TestB(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.B())
	m.SR = 0b00010000
	require.True(t, m.B())
}

func TestD(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.D())
	m.SR = 0b00001000
	require.True(t, m.D())
}

func TestI(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.I())
	m.SR = 0b00000100
	require.True(t, m.I())
}

func TestZ(t *testing.T) {
	m := sim.NewModel()
	require.False(t, m.Z())
	m.SR = 0b00000010
	require.True(t, m.Z())
}

func TestC(t *testing.T) {
	m := sim.NewModel()
	require.Equal(t, byte(0x0), m.C())
	m.SR = 0b00000001
	require.Equal(t, byte(0x1), m.C())
}
