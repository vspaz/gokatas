package filemanipulation

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file: %s", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("line: %s", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s\n", err)
	}
}
