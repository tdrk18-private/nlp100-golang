package section1

import "testing"

func TestMapDictionary(t *testing.T) {
	result := MapDictionary("Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.")
	if !equalMap(result, map[string]int{"Be": 4, "C": 6, "B": 5, "Ca": 20, "F": 9, "S": 16, "H": 1, "K": 19, "Al": 13, "Mi": 12, "Ne": 10, "O": 8, "Li": 3, "P": 15, "Si": 14, "Ar": 18, "Na": 11, "N": 7, "Cl": 17, "He": 2}) {
		t.Fatal("failed test")
	}
}

func equalMap(a, b map[string]int) bool  {
	if len(a) != len(b) {
		return false
	}

	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
