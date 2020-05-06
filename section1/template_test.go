package section1

import "testing"

func TestTemplate(t *testing.T) {
	result := Template(12, "気温", 22.4)
	if result != "12日の気温は22.4" {
		t.Fatal("failed test")
	}
}
