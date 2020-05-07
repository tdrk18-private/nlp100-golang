package section1

import "testing"

func TestNGram(t *testing.T) {
	word := NGram(2, "I am an NLPer", Word)
	if !equalStringArray(word, []string{"I am", "am an", "an NLPer"}) {
		t.Fatal("failed test")
	}

	char := NGram(2, "I am an NLPer", Char)
	if !equalStringArray(char, []string{"Ia", "am", "ma", "an", "nN", "NL", "LP", "Pe", "er"}) {
		t.Fatal("failed test")
	}
}
