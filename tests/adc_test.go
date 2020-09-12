package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ADC a
func Test0x6d(t *testing.T) {
	m, _ := seedA(0x6d, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedA(0x6d, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedA(0x6d, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedA(0x6d, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC a,x
func Test0x7d(t *testing.T) {
	m, _ := seedAX(0x7d, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedAX(0x7d, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedAX(0x7d, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedAX(0x7d, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC a,y
func Test0x79(t *testing.T) {
	m, _ := seedAY(0x79, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedAY(0x79, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedAY(0x79, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedAY(0x79, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC immediate
func Test0x69(t *testing.T) {
	m := raw(0x69, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m = raw(0x69, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m = raw(0x69, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m = raw(0x69, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC zp
func Test0x65(t *testing.T) {
	m, _ := seedZP(0x65, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedZP(0x65, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedZP(0x65, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedZP(0x65, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC (zp,x)
func Test0x61(t *testing.T) {
	m, _ := seedZPIX(0x61, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedZPIX(0x61, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedZPIX(0x61, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedZPIX(0x61, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC zp,x
func Test0x75(t *testing.T) {
	m, _ := seedZPX(0x75, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedZPX(0x75, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedZPX(0x75, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedZPX(0x75, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}

// ADC (zp),y
func Test0x71(t *testing.T) {
	m, _ := seedZPIY(0x71, 0x0c)
	m.A = 0x0f
	m.Tick()
	require.Equal(t, byte(0x1b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.False(t, m.V())

	m, _ = seedZPIY(0x71, 0x0c)
	m.A = 0x7f
	m.Tick()
	require.Equal(t, byte(0x8b), m.A)
	require.True(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(0), m.C())
	require.True(t, m.V())

	m, _ = seedZPIY(0x71, 0x0c)
	m.A = 0xff
	m.Tick()
	require.Equal(t, byte(0x0b), m.A)
	require.False(t, m.N())
	require.False(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())

	m, _ = seedZPIY(0x71, 0x0c)
	m.A = 0xf4
	m.Tick()
	require.Equal(t, byte(0x00), m.A)
	require.False(t, m.N())
	require.True(t, m.Z())
	require.Equal(t, byte(1), m.C())
	require.True(t, m.V())
}
