package concurrency

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func startWorker(timeInterval int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(timeInterval) * time.Second)
	wg.Done()
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	start := time.Now().Second()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go startWorker(5, &wg)
	}
	wg.Wait()
	end := time.Now().Second()
	assert.GreaterOrEqual(t, time.Duration(end-start)*time.Second, 5*time.Second)
	assert.Less(t, time.Duration(end-start)*time.Second, 10*time.Second)
}
