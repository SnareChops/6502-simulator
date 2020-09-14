package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// SEC
func Test0x38(t *testing.T) {
	m := raw(0x38)
	require.Equal(t, byte(0), m.C())
	m.Tick()
	require.Equal(t, byte(1), m.C())
}
