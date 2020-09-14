package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CLC
func Test0x18(t *testing.T) {
	m := raw(0x18)
	m.SEC()
	require.Equal(t, byte(1), m.C())
	m.Tick()
	require.Equal(t, byte(0), m.C())
}
