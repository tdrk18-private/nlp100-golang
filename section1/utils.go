package section1

func equalArray(a, b []int) bool  {
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

func equalMap(a, b map[string]int) bool  {
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

func contains(s []int, t int) bool  {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
