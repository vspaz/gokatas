package strings

func countChars(text string) map[string]int {
	charToCount := map[string]int{}
	for _, char := range text {
		letter := string(char)
		count, _ := charToCount[letter]
		charToCount[letter] = count + 1
	}
	return charToCount
}
