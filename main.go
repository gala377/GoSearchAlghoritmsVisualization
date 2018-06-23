package main

import (
	"log"
	"runtime"

	"github.com/gala377/SearchAlghorithms/alghoritms"
	"github.com/gala377/SearchAlghorithms/graphics"
	"github.com/gala377/SearchAlghorithms/labirynth"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	renderer, win := initGraphicsLibrary()
	defer renderer.Terminate()
	start := false
	renderer.SubscribeKey(graphics.Key(glfw.KeyS), func(state graphics.KeyState) {
		if state == graphics.CLICKED {
			log.Printf("Let the search start")
			start = !start
		}
	})
	algs := initAlghorithms(renderer)
	renderer.SubscribeKey(graphics.Key(glfw.KeyR), func(state graphics.KeyState) {
		if state == graphics.CLICKED {
			start = false
			log.Printf("RESET")
			renderer.Clear()
			algs = initAlghorithms(renderer)
			runtime.GC()
		}
	})
	for !win.ShouldClose() {
		renderer.Draw()
		if start {
			for _, alg := range algs {
				if !alg.Finished() {
					alg.Step()
				}
			}
		}
	}
}

func initGraphicsLibrary() (*graphics.Renderer, *glfw.Window) {
	renderer, err := graphics.New()
	if err != nil {
		panic(err)
	}
	renderer.SetWindowTitle("TestWindow!")
	renderer.SetWindowSize(900, 900)
	win, err := renderer.GetWindow()
	if err != nil {
		panic(err)
	}
	renderer.ConnectCallbacks()
	return renderer, win
}


func initAlghorithms(renderer *graphics.Renderer) []alghoritms.Algorithm {
	boards := initBoards(renderer)
	res := []alghoritms.Algorithm{
		alghoritms.NewBFS(boards[0]),
		alghoritms.NewDFS(boards[1]),
		alghoritms.NewAStar(boards[2]),
	}
	res[2].Init()
	res[0].Init()
	res[1].Init()
	return res
}

func initBoards(renderer *graphics.Renderer) []*alghoritms.SearchSpace {
	lab := labyrinth.Random(50, 50)
	set := labyrinth.DefaultSettings()
	set.Shift = 22
	set.SquareSize = 0.5

	spaces := make([]*alghoritms.SearchSpace, 0)

	b1 := labyrinth.NewBoard(&lab, set)
	b1.Translate(-2.65, -2.7, 0.0)
	renderer.AddObject(b1)
	spaces = append(spaces, alghoritms.NewSpace(b1))

	b2 := labyrinth.NewBoard(&lab, set)
	b2.Translate(0.35, -2.7, 0.0)
	renderer.AddObject(b2)
	spaces = append(spaces, alghoritms.NewSpace(b2))

	b3 := labyrinth.NewBoard(&lab, set)
	b3.Translate(-2.65, 0.3, 0.0)
	renderer.AddObject(b3)
	spaces = append(spaces, alghoritms.NewSpace(b3))

	b4 := labyrinth.NewBoard(&lab, set)
	b4.Translate(0.35, 0.3, 0.0)
	renderer.AddObject(b4)
	spaces = append(spaces, alghoritms.NewSpace(b4))

	return spaces
}
