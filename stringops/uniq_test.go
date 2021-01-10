package stringops

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUniqWithSortedElementsOK(t *testing.T) {
	sortedElements := "1\n2\n3\n4\n5\n6"
	in := bufio.NewReader(strings.NewReader(sortedElements))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	assert.Equal(t, nil, err, "failed")
}
