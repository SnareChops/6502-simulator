package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BMI
func Test0x30(t *testing.T) {
	m := raw(0x30, 0xff)
	m.SetNegative()
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)
	m.Tick()
	require.Equal(t, uint16(0xffff), m.PC)

	m = raw(0x30, 2)
	m.SetNegative()
	m.Tick()
	require.Equal(t, uint16(2), m.PC)

	m = raw(0x30, 2)
	m.Tick()
	require.Equal(t, uint16(1), m.PC)
}
