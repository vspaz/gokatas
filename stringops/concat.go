package stringops

import (
	"strings"
)

func Concat(sep string, words...string) string {
	return strings.Join(words, sep)
}
