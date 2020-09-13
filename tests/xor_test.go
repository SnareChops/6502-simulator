package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// XOR a
func Test0x4d(t *testing.T) {
	m, _ := seedA(0x4d, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0x4d, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR a,x
func Test0x5d(t *testing.T) {
	m, _ := seedAX(0x5d, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAX(0x5d, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR a,y
func Test0x59(t *testing.T) {
	m, _ := seedAY(0x59, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAY(0x59, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR immediate
func Test0x49(t *testing.T) {
	m := raw(0x49, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x49, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR zp
func Test0x45(t *testing.T) {
	m, _ := seedZP(0x45, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0x45, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR (zp,x)
func Test0x41(t *testing.T) {
	m, _ := seedZPIX(0x41, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIX(0x41, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR zp,x
func Test0x55(t *testing.T) {
	m, _ := seedZPX(0x55, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPX(0x55, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// XOR (zp),y
func Test0x51(t *testing.T) {
	m, _ := seedZPIY(0x51, 0b1)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIY(0x51, 0b10000010)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b10000011), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}
