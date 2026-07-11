package main

import (
	"fmt"

	"proxy"
)

func main() {
	fmt.Println("Application Started. Initializing image proxies for gallery...")

	image1 := proxy.NewImageProxy("photo1.jpg")
	_ = proxy.NewImageProxy("photo2.png")
	image3 := proxy.NewImageProxy("photo3.gif")

	fmt.Println("\nGallery initialized. No images actually loaded yet.")
	fmt.Printf("Image 1 Filename: %s\n", image1.GetFileName())

	fmt.Printf("\nUser requests to display %s\n", image1.GetFileName())
	image1.Display()

	fmt.Printf("\nUser requests to display %s again.\n", image1.GetFileName())
	image1.Display()

	fmt.Printf("\nUser requests to display %s\n", image3.GetFileName())
	image3.Display()

	fmt.Println("\nApplication finished. Note: photo2.png was never loaded.")
}
