package stringops

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatWithSeparator(t *testing.T) {
	assert.Equal(
		t,
		"foo.bar.baz",
		Concat(".", "foo", "bar", "baz"),
	)
}
