package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// SBC a
func Test0xed(t *testing.T) {
	m, _ := seedA(0xed, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedA(0xed, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedA(0xed, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedA(0xed, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC a,x
func Test0xfd(t *testing.T) {
	m, _ := seedAX(0xfd, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedAX(0xfd, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedAX(0xfd, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedAX(0xfd, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC a,y
func Test0xf9(t *testing.T) {
	m, _ := seedAY(0xf9, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedAY(0xf9, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedAY(0xf9, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedAY(0xf9, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC immediate
func Test0xe9(t *testing.T) {
	m := raw(0xe9, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m = raw(0xe9, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m = raw(0xe9, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m = raw(0xe9, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC zp
func Test0xe5(t *testing.T) {
	m, _ := seedZP(0xe5, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZP(0xe5, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZP(0xe5, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedZP(0xe5, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC (zp,x)
func Test0xe1(t *testing.T) {
	m, _ := seedZPIX(0xe1, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPIX(0xe1, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPIX(0xe1, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedZPIX(0xe1, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC zp,x
func Test0xf5(t *testing.T) {
	m, _ := seedZPX(0xf5, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPX(0xf5, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPX(0xf5, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedZPX(0xf5, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}

// SBC (zp),y
func Test0xf1(t *testing.T) {
	m, _ := seedZPIY(0xf1, 0x02)
	m.A = 0x0c
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0a), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPIY(0xf1, 0x02)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0x0), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())

	m, _ = seedZPIY(0xf1, 0x04)
	m.A = 0x02
	m.SEC()
	m.Tick()
	require.Equal(t, byte(0xfe), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x00), m.C())
	require.True(t, m.V())

	m, _ = seedZPIY(0xf1, 0x01)
	m.A = 0x10
	m.CLC()
	m.Tick()
	require.Equal(t, byte(0x0e), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0x01), m.C())
	require.False(t, m.V())
}
