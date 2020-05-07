package section1

import (
	"bytes"
	"unicode"
)

func Cipher(x string) string {
	var buf bytes.Buffer

	for i := 0; i < len(x); i++ {
		if isLower(string(x[i])) {
			buf.WriteByte(219 - x[i])
		} else {
			buf.WriteByte(x[i])
		}
	}

	return buf.String()
}

func isLower(x string) bool {
	if len(x) == 0 {
		return false
	}

	for _, s := range []rune(x) {
		if unicode.IsDigit(s) {
			return false
		}
		if unicode.IsUpper(s) {
			return false
		}
	}

	return true
}
