package main

import (
	"github.com/gala377/SearchAlghorithms/graphics"
	"github.com/gala377/SearchAlghorithms/graphics/objects"
)

func main() {
	renderer, err := graphics.New()
	if err != nil {
		panic(err)
	}
	defer renderer.Terminate()
	renderer.SetWindowTitle("TestWindow!")
	renderer.SetWindowSize(800, 600)
	win, err := renderer.GetWindow()
	renderer.ConnectCallbacks()

	renderer.AddObject(objects.NewTriangle2D())
	renderer.AddObject(objects.NewSquare2D(1.0))

	if err != nil {
		panic(err)
	}
	for !win.ShouldClose() {
		renderer.Draw()
	}
}

