package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// STY a
func Test0x8c(t *testing.T) {
	m, a := seedA(0x8c, 0x0)
	m.Y = 0x02
	require.Equal(t, byte(0x0), m.Fetch(a...))
	m.Tick()
	require.Equal(t, byte(0x02), m.Fetch(a...))
}

// STY zp
func Test0x84(t *testing.T) {
	m, a := seedZP(0x84, 0x0)
	m.Y = 0x02
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x02), m.Fetch(a))
}

// STY zp,x
func Test0x94(t *testing.T) {
	m, a := seedZPX(0x94, 0x0)
	m.Y = 0x02
	require.Equal(t, byte(0x0), m.Fetch(a))
	m.Tick()
	require.Equal(t, byte(0x02), m.Fetch(a))
}
