package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseDigitOk(t *testing.T) {
	assert.Equal(t, 4321, ReverseDigit(1234))
	assert.Equal(t, 1, ReverseDigit(1))
	assert.Equal(t, 12, ReverseDigit(21))
}
