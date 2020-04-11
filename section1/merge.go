package section1

import "math"

func Merge(s string, t string) string  {
	runesS := []rune(s)
	runesT := []rune(t)
	var strings []rune
	length := int(math.Max(float64(len(runesS)), float64(len(runesT))))
	for i := 0; i < length; i++ {
		if i < len(runesS) {
			strings = append(strings, runesS[i])
		}
		if i < len(runesT) {
			strings = append(strings, runesT[i])
		}
	}
	return string(strings)
}
