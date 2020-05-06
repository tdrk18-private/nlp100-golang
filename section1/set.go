package section1

import "github.com/golang-collections/collections/set"

func Union(x, y []string) *set.Set {
	setX := createSet(x)
	setY := createSet(y)
	return setX.Union(setY)
}

func Intersection(x, y []string) *set.Set {
	setX := createSet(x)
	setY := createSet(y)
	return setX.Intersection(setY)
}

func Difference(x, y []string) *set.Set {
	setX := createSet(x)
	setY := createSet(y)
	return setX.Difference(setY)
}

func createSet(x []string) *set.Set {
	s := set.New()
	for _, e := range x {
		s.Insert(e)
	}
	return s
}