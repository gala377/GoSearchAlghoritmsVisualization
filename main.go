package main

import (
	"github.com/gala377/SearchAlghorithms/graphics"
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
	if err != nil {
		panic(err)
	}
	for !win.ShouldClose() {
		renderer.Draw()
	}
}

