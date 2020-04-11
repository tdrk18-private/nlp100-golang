package section1

func Split(s string) string  {
	runes := []rune(s)
	var strings []rune
	for i := 0; i < len(runes) - 1; i += 2 {
		strings = append(strings, runes[i])
	}
	return string(strings)
}
