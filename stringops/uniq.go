package stringops

import (
	"bufio"
	"fmt"
	"io"
)


func Uniq(in io.Reader, out io.Writer) error {
	reader := bufio.NewScanner(in)
	var prev string
	var current string
	for reader.Scan() {
		current = reader.Text()
        if current < prev {
        	return fmt.Errorf("unsorted file")
		}
		if current == prev {
			continue
		}
		prev = current
		fmt.Println(out, current)
	}
	return nil
}
