package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBiggestNumInArray(t *testing.T) {
	nums := []int{5, 3, 10, 2, -5, 8, 140, 80, 11}
	assert.Equal(t, 140, getBiggestNumInArray(nums))
}
