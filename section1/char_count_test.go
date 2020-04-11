package section1

import (
	"testing"
)

func TestCharCount(t *testing.T) {
	result := CharCount("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	if !equalArray(result, []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}) {
		t.Fatal("failed test")
	}
}
