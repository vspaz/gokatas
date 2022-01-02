package math

func ReverseDigit(num int) int {
	reversedNum := 0
	for num > 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}
	return reversedNum
}
