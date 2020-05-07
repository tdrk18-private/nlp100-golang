package section1

import (
	"strings"
)

type NGramType int

const (
	Word NGramType = iota
	Char
)

func NGram(n int, s string, t NGramType) []string  {
	var res, slice []string
	var delimiter string

	switch t {
	case Word:
		slice = strings.Split(s, " ")
		delimiter = " "
	case Char:
		s = strings.Replace(s, " ", "", len(s))
		slice = strings.Split(s, "")
		delimiter = ""
	}

	length := len(slice)
	for i := 0; i < length - n + 1; i++ {
		res = append(res, strings.Join(slice[i:i+n], delimiter))
	}

	return res
}
