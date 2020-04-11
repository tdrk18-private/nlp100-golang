package section1

import "testing"

func TestEqualArray(t *testing.T) {
	result := equalArray([]int{1, 2, 3}, []int{1, 2, 3})
	if !result {
		t.Fatal("failed test")
	}
	result = equalArray([]int{1, 2, 3}, []int{1, 2, 4})
	if result {
		t.Fatal("failed test")
	}
	result = equalArray([]int{1, 2, 3}, []int{1, 2})
	if result {
		t.Fatal("failed test")
	}
}

func TestEqualMap(t *testing.T) {
	result := equalMap(map[string]int{"hoge": 1, "fuga": 2, "bar": 3}, map[string]int{"hoge": 1, "fuga": 2, "bar": 3})
	if !result {
		t.Fatal("failed test")
	}
	result = equalMap(map[string]int{"hoge": 1, "fuga": 2, "bar": 3}, map[string]int{"hoge": 1, "fuga": 2, "bar": 4})
	if result {
		t.Fatal("failed test")
	}
	result = equalMap(map[string]int{"hoge": 1, "fuga": 2, "bar": 3}, map[string]int{"hoge": 1, "fuga": 2})
	if result {
		t.Fatal("failed test")
	}
}

func TestContains(t *testing.T) {
	result := contains([]int{1, 2, 3}, 1)
	if !result {
		t.Fatal("failed test")
	}
	result = contains([]int{1, 2, 3}, 5)
	if result {
		t.Fatal("failed test")
	}
}