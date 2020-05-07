package section1

import (
	"strings"
	"testing"
)

func TestTypoglycemia(t *testing.T) {
	text := "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	result := Typoglycemia(text)
	list := strings.Split(result, " ")
	count := len(list)
	if count != 21 {
		t.Fatal("test failed")
	}
	if list[0] != "I" {
		t.Fatal("test failed")
	}
	if list[3] != "that" {
		t.Fatal("test failed")
	}
	if list[8] != "what" {
		t.Fatal("test failed")
	}
	if list[10] != "was" {
		t.Fatal("test failed")
	}
	if list[19] != "mind" {
		t.Fatal("test failed")
	}
	if !(list[2][0] == 'b' && list[2][len(list[2])-1] == 'e') {
		t.Fatal("test failed")
	}
	if !(list[5][0] == 'c' && list[5][len(list[5])-1] == 'd') {
		t.Fatal("test failed")
	}
	if !(list[6][0] == 'a' && list[6][len(list[6])-1] == 'y') {
		t.Fatal("test failed")
	}
	if !(list[7][0] == 'u' && list[7][len(list[7])-1] == 'd') {
		t.Fatal("test failed")
	}
	if !(list[18][0] == 'h' && list[18][len(list[18])-1] == 'n') {
		t.Fatal("test failed")
	}
}
