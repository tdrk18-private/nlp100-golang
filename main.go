package main

import (
	"./section1"
	"fmt"
)

func main() {
	println(section1.Reverse("stressed"))
	println(section1.Split("パタトクカシーー"))
	println(section1.Merge("パトカー", "タクシー"))
	fmt.Printf("%d\n", section1.CharCount("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."))
	fmt.Printf("%v\n", section1.MapDictionary("Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."))
}
