package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// LDA a
func Test0xad(t *testing.T) {
	m, _ := seedA(0xad, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xad, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xad, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA a,x
func Test0xbd(t *testing.T) {
	m, _ := seedAX(0xbd, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0xbd, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAX(0xbd, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA a,y
func Test0xb9(t *testing.T) {
	m, _ := seedAY(0xb9, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0xb9, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAY(0xb9, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA immediate
func Test0xa9(t *testing.T) {
	m := raw(0xa9, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.N())

	m = raw(0xa9, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xa9, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA zp
func Test0xa5(t *testing.T) {
	m, _ := seedZP(0xa5, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xa5, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xa5, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA (zp,x)
func Test0xa1(t *testing.T) {
	m, _ := seedZPIX(0xa1, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIX(0xa1, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIX(0xa1, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA zp,x
func Test0xb5(t *testing.T) {
	m, _ := seedZPX(0xb5, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0xb5, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPX(0xb5, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}

// LDA (zp),y
func Test0xb1(t *testing.T) {
	m, _ := seedZPIY(0xb1, 0x0c)
	m.Tick()
	require.Equal(t, byte(0x0c), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIY(0xb1, 0x0)
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIY(0xb1, 0x82)
	m.Tick()
	require.Equal(t, byte(0x82), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
}
