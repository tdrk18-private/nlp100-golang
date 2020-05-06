package section1

import "testing"

func TestCipher(t *testing.T) {
	x := "re1Ae"
	encrypt := Cipher(x)
	if encrypt != "iv1Av" {
		t.Fatal("failed test")
	}
	decrypt := Cipher(encrypt)
	if decrypt != "re1Ae" {
		t.Fatal("failed test")
	}
}
