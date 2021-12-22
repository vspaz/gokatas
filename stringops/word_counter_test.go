package stringops

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountWordsOk(t *testing.T) {
	testInput := bytes.NewBufferString("word1 word2 word3\n")
	assert.Equal(t, 3, CountWords(testInput))
}
