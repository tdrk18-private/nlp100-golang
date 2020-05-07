package section1

import (
	"fmt"
	"math/rand"
	"strings"
)

func Typoglycemia(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for _, l := range list {
		if len(l) <= 4 {
			res = append(res, l)
		} else {
			head := string(l[0])
			tail := string(l[len(l)-1])
			internal := strings.Split(l[1:len(l)-1], "")
			shuffle(internal)
			word := strings.Join(internal, "")
			res = append(res, fmt.Sprintf("%s%s%s", head, word, tail))
		}
	}
	return strings.Join(res, " ")
}

func shuffle(s []string) {
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
