package main

import (
	"fmt"

	"flyweight"
)

func main() {
	editor := flyweight.NewTextEditorClient()

	fmt.Println("Rendering text with same properties:")
	editor.RenderText("Hello", 0, 0, "Arial", 12, "black")
	fmt.Printf("Number of unique characters: %d\n\n", editor.GetUniqueCharacterCount())

	fmt.Println("Rendering text with different properties:")
	editor.RenderText("Hello", 0, 20, "Times New Roman", 14, "blue")
	fmt.Printf("Number of unique characters: %d\n\n", editor.GetUniqueCharacterCount())

	fmt.Println("Rendering text with mixed properties:")
	editor.RenderText("World", 0, 40, "Arial", 12, "red")
	fmt.Printf("Number of unique characters: %d\n", editor.GetUniqueCharacterCount())
}
