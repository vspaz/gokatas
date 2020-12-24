package main

import "fmt"

func main() {
	someString := "foobarbazzzzfufufuf"
	charToCount := map[int32]int{}
	for _, char := range(someString) {
		count, _ := charToCount[char]
		charToCount[char] = count + 1
	}
	for char, count := range(charToCount) {
        fmt.Printf("%c => %d\n", char, count)
	}
}
