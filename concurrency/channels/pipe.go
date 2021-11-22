package channels

func generateValuesAndClosePipe(pipe chan int) {
	for i := 0; i < 10; i++ {
		pipe <- i
	}
	close(pipe)
}

func generateValues(pipe chan int) {
	defer close(pipe)
	for i := 0; i < 10; i++ {
		pipe <- i
	}
}
