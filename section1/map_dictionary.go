package section1

import "strings"

func MapDictionary(s string) map[string]int  {
	slice := strings.Split(s, " ")
	var maps = make(map[string]int)
	for i, _ := range slice {
		text := slice[i]
		if contains([]int{0, 4, 5, 6, 7, 8, 14, 15, 18}, i) {
			maps[text[:1]] = i + 1
		} else {
			maps[text[:2]] = i + 1
		}
	}
	return maps
}

func contains(s []int, t int) bool  {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}