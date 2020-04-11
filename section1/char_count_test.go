package section1

import (
	"testing"
)

func TestCharCount(t *testing.T) {
	result := CharCount("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	if !equal(result, []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}) {
		t.Fatal("failed test")
	}
}

func equal(a, b []int) bool  {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
