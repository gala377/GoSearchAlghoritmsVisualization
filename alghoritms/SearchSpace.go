package alghoritms

import (
	"github.com/gala377/SearchAlghorithms/labirynth"
)

var ACTIVE_COLOR = struct{R, G, B float32}{1.0, 0.0, 0.0}
var VISITED_COLOR = struct{R, G, B float32}{0.5, 0.0, 0.0}
var FRONTIRE_COLOR = struct{R, G, B float32}{0.9, 0.6, 0.0}


type Point struct{ X, Y uint }

type SearchSpace struct {
	*labyrinth.Board
	active []Point
}

func NewSpace(board *labyrinth.Board) *SearchSpace {
	return &SearchSpace{board, make([]Point, 0)}
}

func (space* SearchSpace) Get(x, y uint) bool {
	return space.Board.Labyrinth.GetField(x, y)
}

func (space* SearchSpace) Visit(x, y uint) {
	//log.Printf("Visiting space (%d, %d)", x, y)
	space.Board.SetColor(x, y,
		ACTIVE_COLOR.R, ACTIVE_COLOR.G, ACTIVE_COLOR.B)
	//log.Printf("Visited")
	space.active = append(space.active, struct{X, Y uint}{x, y})
}

func (space* SearchSpace) MarkFrontier(frontier ...Point) {
	for _, el := range frontier {
		//log.Printf("Visiting space (%d, %d)", x, y)
		space.Board.SetColor(el.X, el.Y,
			FRONTIRE_COLOR.R, FRONTIRE_COLOR.G, FRONTIRE_COLOR.B)
		//log.Printf("Visited")
		space.active = append(space.active, el)
	}
}

func (space* SearchSpace) Update() {
	for _, point := range space.active {
		space.Board.SetColor(point.X, point.Y,
			VISITED_COLOR.R, VISITED_COLOR.G, VISITED_COLOR.B)
	}
	space.active = make([]Point, 0)
}
