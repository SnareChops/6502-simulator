package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CMP a
func Test0xcd(t *testing.T) {
	m, _ := seedA(0xcd, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedA(0xcd, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedA(0xcd, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP a,x
func Test0xdd(t *testing.T) {
	m, _ := seedAX(0xdd, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAX(0xdd, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAX(0xdd, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP a,y
func Test0xd9(t *testing.T) {
	m, _ := seedAY(0xd9, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedAY(0xd9, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedAY(0xd9, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP immediate
func Test0xc9(t *testing.T) {
	m := raw(0xc9, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m = raw(0xc9, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m = raw(0xc9, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP zp
func Test0xc5(t *testing.T) {
	m, _ := seedZP(0xc5, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZP(0xc5, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZP(0xc5, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP (zp,x)
func Test0xc1(t *testing.T) {
	m, _ := seedZPIX(0xc1, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIX(0xc1, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIX(0xc1, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP zp,x
func Test0xd5(t *testing.T) {
	m, _ := seedZPX(0xd5, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPX(0xd5, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPX(0xd5, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}

// CMP (zp),y
func Test0xd1(t *testing.T) {
	m, _ := seedZPIY(0xd5, 1)
	m.A = 0
	m.Tick()
	require.Equal(t, byte(0), m.C())
	require.True(t, m.N())
	require.False(t, m.Z())

	m, _ = seedZPIY(0xd1, 1)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.True(t, m.Z())

	m, _ = seedZPIY(0xd1, 0)
	m.A = 1
	m.Tick()
	require.Equal(t, byte(1), m.C())
	require.False(t, m.N())
	require.False(t, m.Z())
}
