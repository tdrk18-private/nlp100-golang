package section1

import "github.com/golang-collections/collections/set"

func equalArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func equalSet(a, b *set.Set) bool  {
	if a.Len() != b.Len() {
		return false
	}

	return a.SubsetOf(b) && b.SubsetOf(a)
}

func contains(s []int, t int) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
