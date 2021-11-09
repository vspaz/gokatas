package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func worker(id int, channel chan string, waitInterval time.Duration) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("worker %d; message '%d'", id, i)
		time.Sleep(waitInterval)
	}
}

func TestSelect(t *testing.T) {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go worker(1, channel_1, time.Millisecond*100)
	go worker(2, channel_2, time.Second*5)

	for i := 0; i < 5; i++ {
		assert.Equal(t, fmt.Sprintf("worker 1; message '%d'", i), <-channel_1)
		assert.Equal(t, fmt.Sprintf("worker 2; message '%d'", i), <-channel_2)  // blocking
	}
}
