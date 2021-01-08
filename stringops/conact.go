package stringops

import (
	"strings"
)

func Concat(words ...string) string {
	return strings.Join(words, " ")
}
