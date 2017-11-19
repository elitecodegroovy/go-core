package util

import (
	"bytes"
	"unicode"
)

// Swap letter to upper or lower.
func SwapCase(str string) string {
	buf := &bytes.Buffer{}
	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(unicode.ToUpper(r))
		}
	}
	return buf.String()
}
