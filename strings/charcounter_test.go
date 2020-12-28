package strings

import "testing"
import "reflect"

func TestCountChars(t *testing.T) {
	got := countChars("aabbcc")
	expected := map[int32]int{97: 2, 98: 2, 99: 2}
	reflect.DeepEqual(expected, got)
}