package stringops

import (
	"bufio"
	"io"
)

func countWords(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var wordCount int

	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}
