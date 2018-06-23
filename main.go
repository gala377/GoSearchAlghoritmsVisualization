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
	renderer.SetWindowSize(1280, 720)
	win, err := renderer.GetWindow()
	renderer.ConnectCallbacks()

	lab := labyrinth.Random(50, 50)
	set := labyrinth.DefaultSettings()
	set.Shift = 20
	set.SquareSize = 0.5
	board := labyrinth.NewBoard(lab, set)
	board.Translate(-2.0, -2.0, 0.0)
	renderer.AddObject(board)

	if err != nil {
		panic(err)
	}
	for !win.ShouldClose() {
		renderer.Draw()
	}
}
