package stringops

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestUniqWithSortedElementsOK(t *testing.T) {
	sortedElements := "1\n2\n3\n4\n5\n6"
	in := bufio.NewReader(strings.NewReader(sortedElements))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	assert.Equal(t, nil, err, "failed")
}