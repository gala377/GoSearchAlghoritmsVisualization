package main

import (
	"github.com/gala377/SearchAlghorithms/graphics"
	"github.com/gala377/SearchAlghorithms/labirynth"
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

	lab := labyrinth.Random(100, 100)
	board := labyrinth.NewBoard(lab)
	board.Translate(-3.0, -2.9, 0.0)
	renderer.AddObject(&board)

	if err != nil {
		panic(err)
	}
	for !win.ShouldClose() {
		renderer.Draw()
	}
}
