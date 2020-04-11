package main

import (
	"./section1"
)

func main() {
	println(section1.Reverse("stressed"))
	println(section1.Split("パタトクカシーー"))
	println(section1.Merge("パトカー", "タクシー"))
}
