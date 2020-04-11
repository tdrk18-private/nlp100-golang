package section1

import "testing"

func TestReverse(t *testing.T) {
	result := Reverse("stressed")
	if result != "desserts" {
		t.Fatal("failed test")
	}
}
