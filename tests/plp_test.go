package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0x28(t *testing.T) {
	m := raw(0x28)
	m.Push(0b11001111)
	m.Tick()
	require.Equal(t, byte(0b11101111), m.SR)
}
