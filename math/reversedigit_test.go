package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseDigit(t *testing.T) {
	assert.Equal(t, 4321, ReverseDigit(1234))
}
