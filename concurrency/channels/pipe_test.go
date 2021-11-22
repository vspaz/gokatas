package channels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPipe(t *testing.T) {
	pipe := make(chan int)
	go generateValuesAndClosePipe(pipe)
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, <-pipe)
	}
}

func TestClosedPipe(t *testing.T) {
	pipe := make(chan int)
	var nums []int
	go generateValuesAndClosePipe(pipe)
	for i := 0; i < 11; i++ {
		nums = append(nums, <-pipe)
	}
	assert.Equal(t, 0, nums[0])
	assert.Equal(t, 0, nums[10]) // pipe is already closed
}

func TestIsPipeClosed(t *testing.T) {
	pipe := make(chan int)
	go generateValuesAndClosePipe(pipe)
	pipeStatus := "open"
	for i := 0; i < 11; i++ {
		_, isOpen := <-pipe
		if !isOpen {
			pipeStatus = "closed"
			break
		}
	}
	assert.Equal(t, "closed", pipeStatus)
}

func TestClosedPipeWithRange(t *testing.T) {
	pipe := make(chan int)
	go generateValuesAndClosePipe(pipe)
	var num int
	for num = range pipe {
	}
	assert.Equal(t, 9, num)
}

func TestDeferClosingPipe(t *testing.T) {
	pipe := make(chan int)
	go generateValues(pipe)
	var nums []int
	for i := 0; i < 11; i++ {
		nums = append(nums, <-pipe)
	}
	assert.Equal(t, 0, nums[0])
	assert.Equal(t, 0, nums[10]) // pipe is already closed
}
