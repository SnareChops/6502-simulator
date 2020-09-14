package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// PHP
func Test0x08(t *testing.T) {
	m := raw(0x08)
	m.SR = 0xff
	m.Tick()
	require.Equal(t, byte(0b11001111), m.Pop())
}
