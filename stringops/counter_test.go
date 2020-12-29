package stringops

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstNonRepeatedCharacterWithOneNonRepeatedCharacter(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("aabbcdd")
	expected := "c"
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}

func TestGetFirstNonRepeatedCharacterWithMultipleNonRepeatedCharacters(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("aabbcddeff")
	expected := "c"
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}

func TestGetFirstNonRepeatedCharacterAllCharactersAreRepeataed(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("aabbccddeeff")
	expected := ""
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}

func TestGetFirstNonRepeatedCharacterEmptyString(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("")
	expected := ""
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}

func TestGetFirstNonRepeatedCharacterAllNonRepeatedCharacters(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("abcdef")
	expected := "a"
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}

func TestGetFirstNonRepeatedCharacterNonRepeatedCharacterComesLast(t *testing.T) {
	got := GetFirstNonRepeatedCharacter("aabbccddeef")
	expected := "f"
	assert.Equal(t, got, expected, "expected '%v'; got '%v'", got, expected)
}
