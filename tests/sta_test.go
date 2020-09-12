package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// STA a
func Test0x8d(t *testing.T) {
	m, a := seedA(0x8d, 0x0)
	m.A = 0x0f
	require.Equal(t, byte(0x0), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a...))
}

// STA a,x
func Test0x9d(t *testing.T) {
	m, a := seedAX(0x9d, 0x0)
	m.A = 0x0f
	require.Equal(t, byte(0x0), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a...))
}

// STA a,y
func Test0x99(t *testing.T) {
	m, a := seedAY(0x99, 0x0)
	m.A = 0x0f
	require.Equal(t, byte(0x0), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a...))
}

// STA zp
func Test0x85(t *testing.T) {
	m, a := seedZP(0x85, 0x0)
	m.A = 0x0f
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a))
}

//STA (zp,x)
func Test0x81(t *testing.T) {
	m, a := seedZPIX(0x81, 0x0c)
	m.A = 0x0f
	require.Equal(t, byte(0x0c), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a...))
}

// STA zp,x
func Test0x95(t *testing.T) {
	m, a := seedZPX(0x95, 0x0)
	m.A = 0x0f
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a))
}

// STA (zp),y
func Test0x91(t *testing.T) {
	m, a := seedZPIY(0x91, 0x0c)
	m.A = 0x0f
	require.Equal(t, byte(0x0c), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x0f), m.Fetch(a...))
}
