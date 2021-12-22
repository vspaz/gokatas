package stringops

import (
	"bufio"
	"io"
)

func CountWords(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var wordCount int

	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}