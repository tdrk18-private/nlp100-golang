package section1

import "testing"

func TestMerge(t *testing.T) {
	result := Merge("パトカー", "タクシー")
	if result != "パタトクカシーー" {
		t.Fatal("failed test")
	}
}
