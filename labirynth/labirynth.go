package labyrinth

import (
	"log"
)

type Labyrinth struct {
	size_x uint
	size_y uint
	board  [][]bool
}

func (l Labyrinth) GetField(x, y uint) bool {
	return l.board[x][y]
}

func (l Labyrinth) Dimensions() (x, y uint) {
	return l.size_x, l.size_y
}

func Random(size_x, size_y uint) Labyrinth {
	board := make([][]bool, size_x)
	for i := 0; uint(i) < size_x; i++ {
		board[i] = make([]bool, size_y)
	}
	log.Println("Making Random labyrnthm")
	var temp string
	for _, column := range board {
		for _ = range column {
			temp += "|x|"
		}
		temp += "\n"
	}
	log.Println(temp)
	return Labyrinth{size_x: size_x, size_y: size_y, board: board}
}
