package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CLV
func Test0xb8(t *testing.T) {
	m := raw(0xb8)
	m.SetOverflow()
	require.True(t, m.V())
	m.Tick()
	require.False(t, m.V())
}
