package main

import (
	"log"

	"github.com/gala377/SearchAlghorithms/graphics"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	log.Println("Initializing program")
	renderer, err := graphics.New()
	if err != nil {
		panic(err)
	}
	defer renderer.Terminate()
	renderer.SetWindowTitle("TestWindow!")
	renderer.SetWindowSize(800, 600)
	win, err := renderer.GetWindow()
	if err != nil {
		panic(err)
	}
	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.PollEvents()
	}
}

