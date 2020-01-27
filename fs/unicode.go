package fs

import (
	"strconv"
	"strings"
)

func Unicode(s string) string {
	r, _ := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)

	return string(r)
}
