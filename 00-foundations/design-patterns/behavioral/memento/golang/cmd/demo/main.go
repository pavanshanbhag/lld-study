package main

import (
	"fmt"

	"memento"
)

func main() {
	editor := memento.NewTextEditor()
	undoManager := memento.NewTextEditorUndoManager()

	editor.Type("Hello, ")
	undoManager.Save(editor)

	editor.Type("world!")
	undoManager.Save(editor)

	editor.Type(" This is a test.")
	fmt.Println("Current content:", editor.GetContent())

	undoManager.Undo(editor)
	fmt.Println("After undo:", editor.GetContent())

	undoManager.Undo(editor)
	fmt.Println("After second undo:", editor.GetContent())

	undoManager.Undo(editor)
}
