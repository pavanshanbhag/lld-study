package main

import (
	"fmt"

	"composite"
)

func main() {
	root := composite.NewFolder("Root")

	documents := composite.NewFolder("Documents")
	downloads := composite.NewFolder("Downloads")
	pictures := composite.NewFolder("Pictures")

	report := composite.NewFile("report.txt", 1024)
	image := composite.NewFile("image.jpg", 2048)
	video := composite.NewFile("video.mp4", 4096)

	root.Add(documents)
	root.Add(downloads)
	root.Add(pictures)

	documents.Add(report)
	downloads.Add(video)
	pictures.Add(image)

	fmt.Println("File System Structure:")
	root.Print("")

	fmt.Printf("\nTotal size: %d bytes\n", root.GetSize())

	fmt.Println("\nDeleting Documents folder:")
	root.Remove(documents)
	documents.Delete()

	fmt.Println("\nUpdated File System Structure:")
	root.Print("")
}
