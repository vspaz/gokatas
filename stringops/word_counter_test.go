package stringops

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountWordsOk(t *testing.T) {
	testInput := bytes.NewBufferString("word1 word2 word3\n")
	assert.Equal(t, 3, count(testInput, false))
}

func TestCountLinesOk(t *testing.T) {
	testInput := bytes.NewBufferString("word1\n word2\n word3\n\n")
	assert.Equal(t, 4, count(testInput, true))
}
