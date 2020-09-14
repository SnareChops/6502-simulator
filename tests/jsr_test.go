package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0x20(t *testing.T) {
	m := raw(0x20, 0xaa, 0xbb)
	m.Tick()
	require.Equal(t, byte(2), m.Fetch(m.SP, 0x01))
	// TODO Expand test to check lsb as well
}
