package stringops

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUniqWithSortedElementsOK(t *testing.T) {
	sortedElements := "1\n2\n3\n4\n5\n6\n"
	in := bufio.NewReader(strings.NewReader(sortedElements))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	assert.Equal(t, nil, err, "failed")
}

func TestUniqWithUnsortedElementsFail(t *testing.T) {
	unsortedElements := "1\n10\n\n2\n"
	in := bufio.NewReader(strings.NewReader(unsortedElements))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	assert.NotEqual(t, nil, err, "failed")
}
