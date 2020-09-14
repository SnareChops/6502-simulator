package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// PHA
func Test0x48(t *testing.T) {
	m := raw(0x48)
	m.A = 0x56
	m.Tick()
	require.Equal(t, byte(0x56), m.Fetch(0xfc, 0x01))
	require.Equal(t, byte(0xfc), m.SP)
}
