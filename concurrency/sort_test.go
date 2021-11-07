package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []int{10, 0, -1, 31, 150, -50, 80, 90, 300, 1000, -180}
	assert.Equal(t, []int{-180, -50, -1, 0, 10, 31, 80, 90, 150, 300, 1000}, MergeSort(data))
}
