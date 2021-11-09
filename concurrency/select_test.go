package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel_1 <- fmt.Sprintf("worker 1; message '%s'", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel_2 <- fmt.Sprintf("worker 2; message '%s'", i)
		}
	}()
}
