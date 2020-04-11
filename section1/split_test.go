package section1

import (
	"testing"
)

func TestSplit(t *testing.T) {
	result := Split("パタトクカシーー")
	if result != "パトカー" {
		t.Fatal("failed test")
	}
}
