package math

func getDigitCount(num int) int {
	count := 0
	remainder := num

	for remainder > 0 {
		remainder /= 10
		count++
	}
	return count
}
