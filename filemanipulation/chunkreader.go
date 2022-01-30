package filemanipulation

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s\n", err)
	}
}
