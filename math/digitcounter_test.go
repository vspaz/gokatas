package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDigitCountOk(t *testing.T) {
	assert.Equal(t, 4, getDigitCount(1234))
	assert.Equal(t, 1, getDigitCount(1))
}
