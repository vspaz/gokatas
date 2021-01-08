package stringops

func countChars(text string) map[string]int {
	charToCount := map[string]int{}
	for _, char := range text {
		charToCount[string(char)]++
	}
	return charToCount
}
