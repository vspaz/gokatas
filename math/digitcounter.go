package math

func getDigitCount(num int) int {
	var digitCount int
	remainder := num

	for remainder > 0 {
		remainder = digitCount / 10
		digitCount++
	}
	return digitCount
}
