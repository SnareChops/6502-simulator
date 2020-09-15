package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0x40(t *testing.T) {
	m := raw(0x40)
	m.Push(0xbb)
	m.Push(0xaa)
	m.Push(0b01100101)
	m.Tick()
	require.Equal(t, uint16(0xbbab), m.PC)
	require.Equal(t, byte(0b01100101), m.SR.Byte())
}
