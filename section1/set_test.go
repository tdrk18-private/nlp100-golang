package section1

import (
	"github.com/golang-collections/collections/set"
	"testing"
)

func TestSet(t *testing.T)  {
	x := NGram(2, "paraparaparadise", Char)
	y := NGram(2, "paragraph", Char)
	union := Union(x, y)
	if !equalSet(union, set.New("pa", "ar", "ra", "ap", "ad", "di", "is", "se", "ag", "gr", "ph")) {
		t.Fatal("failed test")
	}
}

func TestIntersection(t *testing.T) {
	x := NGram(2, "paraparaparadise", Char)
	y := NGram(2, "paragraph", Char)
	intersection := Intersection(x, y)
	if !equalSet(intersection, set.New("pa", "ar", "ra", "ap")) {
		t.Fatal("failed test")
	}
}

func TestDifference(t *testing.T) {
	x := NGram(2, "paraparaparadise", Char)
	y := NGram(2, "paragraph", Char)
	difference := Difference(x, y)
	if !equalSet(difference, set.New("ad", "di", "is", "se")) {
		t.Fatal("failed test")
	}
}
