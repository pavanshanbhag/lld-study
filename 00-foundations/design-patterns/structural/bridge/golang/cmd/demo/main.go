package main

import (
	"fmt"

	"design-patterns/golang/bridge"
)

func main() {
	vectorRenderer := bridge.NewVectorRenderer()
	rasterRenderer := bridge.NewRasterRenderer()

	circle := bridge.NewCircle(vectorRenderer, 5.0)
	rectangle := bridge.NewRectangle(vectorRenderer, 10.0, 5.0)

	fmt.Println("Drawing shapes with vector renderer:")
	circle.Draw()
	rectangle.Draw()

	circle = bridge.NewCircle(rasterRenderer, 5.0)
	rectangle = bridge.NewRectangle(rasterRenderer, 10.0, 5.0)

	fmt.Println("\nDrawing shapes with raster renderer:")
	circle.Draw()
	rectangle.Draw()
}
