package stringops

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountChars(t *testing.T) {
	got := countChars("aabbcc")
	expected := map[string]int{"a": 2, "b": 2, "c": 2}
	assert.Equal(t, expected, got)
}
func TestCountCharsEmptyString(t *testing.T) {
	got := countChars("")
	expected := map[string]int{}
	assert.Equal(t, expected, got)
}
func TestCountCharsDifferentCount(t *testing.T) {
	got := countChars("aabbbccccg")
	expected := map[string]int{"a": 2, "b": 3, "c": 4, "g": 1}
	assert.Equal(t, expected, got)
}
