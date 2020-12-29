package strings

import "testing"
import "reflect"

func TestCountChars(t *testing.T) {
	got := countChars("aabbcc")
	expected := map[int32]int{97: 2, 98: 2, 99: 2}
	reflect.DeepEqual(expected, got)
}
func TestCountCharsEmptyString(t *testing.T) {
	got := countChars("")
	expected := map[int32]int{}
	reflect.DeepEqual(expected, got)
}
func TestCountCharsDifferentCount(t *testing.T) {
	got := countChars("aabbcccdddeeeefffffg")
	expected := map[int32]int{97: 2, 98: 3, 99: 4, 100: 5, 101: 1}
	reflect.DeepEqual(expected, got)
}
