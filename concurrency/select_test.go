package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	waitIntervalForWorker_1 = time.Second * 5
	waitIntervalForWorker_2 = time.Millisecond * 100
)

func worker(id int, channel chan string, waitInterval time.Duration) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("worker %d; message '%d'", id, i)
		time.Sleep(waitInterval)
	}
}

func TestGoroutineBlocked(t *testing.T) {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go worker(1, channel_1, time.Millisecond*100)
	go worker(2, channel_2, waitIntervalForWorker_2)

	for i := 0; i < 5; i++ {
		start := time.Now()
		assert.Equal(t, fmt.Sprintf("worker 2; message '%d'", i), <-channel_2)
		end := time.Now()
		assert.GreaterOrEqual(t, waitIntervalForWorker_2, time.Duration(end.Second()-start.Second())) // blocked for waitInterval seconds.
		assert.Equal(t, fmt.Sprintf("worker 1; message '%d'", i), <-channel_1)
	}
}

func TestGoroutineNonBlocking(t *testing.T) {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go worker(1, channel_1, waitIntervalForWorker_1)
	go worker(2, channel_2, waitIntervalForWorker_2)

	var count int
	for i := 0; i < 5; i++ {
		start := time.Now()
		select {
		case _ = <-channel_2:
			count++
		case _ = <-channel_1:
			end := time.Now()
			assert.LessOrEqual(t, time.Duration(end.Second()-start.Second()), waitIntervalForWorker_1) // non-blocking
			assert.Equal(t, i, count)
			count++
		}
	}
}

func TestSelectDefault(t *testing.T) {

}