package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ORA a
func Test0x0d(t *testing.T) {
	m, _ := seedA(0x0d, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0x0d, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0x0d, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA a,x
func Test0x1d(t *testing.T) {
	m, _ := seedAX(0x1d, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0x1d, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAX(0x1d, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA a,y
func Test0x19(t *testing.T) {
	m, _ := seedAY(0x19, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0x19, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAY(0x19, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA immediate
func Test0x09(t *testing.T) {
	m := raw(0x09, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x09, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0x09, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA zp
func Test0x05(t *testing.T) {
	m, _ := seedZP(0x05, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0x05, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0x05, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA (zp,x)
func Test0x01(t *testing.T) {
	m, _ := seedZPIX(0x01, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIX(0x01, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIX(0x01, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA zp,x
func Test0x15(t *testing.T) {
	m, _ := seedZPX(0x15, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0x15, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPX(0x15, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// ORA (zp),y
func Test0x11(t *testing.T) {
	m, _ := seedZPIY(0x11, 0b10)
	m.A = 0b1
	m.Tick()
	require.Equal(t, byte(0b11), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIY(0x11, 0b0)
	m.A = 0b0
	m.Tick()
	require.Equal(t, byte(0b0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIY(0x11, 0b01000000)
	m.A = 0b10000000
	m.Tick()
	require.Equal(t, byte(0b11000000), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}
