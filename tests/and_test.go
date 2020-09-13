package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// AND a
func Test0x2d(t *testing.T) {
	m, _ := seedA(0x2d, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0x2d, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0x2d, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND a,x
func Test0x3d(t *testing.T) {
	m, _ := seedAX(0x3d, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0x3d, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0x3d, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND a,y
func Test0x39(t *testing.T) {
	m, _ := seedAY(0x39, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0x39, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0x39, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND immediate
func Test0x29(t *testing.T) {
	m := raw(0x29, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m = raw(0x29, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0x29, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND zp
func Test0x25(t *testing.T) {
	m, _ := seedZP(0x25, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0x25, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0x25, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND (zp,x)
func Test0x21(t *testing.T) {
	m, _ := seedZPIX(0x21, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIX(0x21, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIX(0x21, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND zp,x
func Test0x35(t *testing.T) {
	m, _ := seedZPX(0x35, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0x35, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0x35, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}

// AND (zp),y
func Test0x31(t *testing.T) {
	m, _ := seedZPIY(0x31, 0b01000101)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIY(0x31, 0xff)
	m.A = 0x80
	m.Tick()
	require.Equal(t, byte(0x80), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIY(0x31, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
}
