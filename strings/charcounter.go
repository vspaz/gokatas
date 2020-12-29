package strings

func countChars(text string) map[int32]int {
	charToCount := map[int32]int{}
	for _, char := range text {
		count, _ := charToCount[char]
		charToCount[char] = count + 1
	}
	return charToCount
}
