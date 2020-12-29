package stringops

import (
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func GetFirstNonRepeatedCharacter(someString string) string {
	charToCount := linkedhashmap.New()
	for _, char := range someString {
		count, ok := charToCount.Get(char)
		if ok {
			charToCount.Put(char, count.(int)+1)
		} else {
			charToCount.Put(char, 1)
		}
	}
	for _, k := range charToCount.Keys() {
		value, _ := charToCount.Get(k)
		if value == 1 {
			return fmt.Sprintf("%c", k)
		}
	}
	return ""
}
