package section1

import "strings"

func CharCount(s string) []int  {
	slice := strings.Split(s, " ")
	rep := strings.NewReplacer(",", "", ".", "")
	var counts []int
	for i, _ := range slice {
		text := rep.Replace(slice[i])
		counts = append(counts, len(text))
	}
	return counts
}
