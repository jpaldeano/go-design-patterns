package main

import "fmt"

// DecorateTree decorates the tree based on the style
func DecorateTree(treeType string) {
	switch treeType {
	case "classic":
		fmt.Println("Decorating the tree with tinsel, baubles, and a star.")
	case "modern":
		fmt.Println("Decorating the tree with monochrome ornaments and LED lights.")
	case "minimalist":
		fmt.Println("Decorating the tree with a few white lights and simple ornaments.")
	default:
		fmt.Println("Unknown tree style. Leaving it undecorated.")
	}
}

func main() {
	DecorateTree("classic")
	DecorateTree("modern")
	DecorateTree("minimalist")
	DecorateTree("vintage") // Unsupported style
}
