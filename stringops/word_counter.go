package stringops

import (
	"bufio"
	"io"
)

func count(reader io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(reader)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	var wordCount int

	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}
