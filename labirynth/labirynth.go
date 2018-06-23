package labyrinth

import (
	"math/rand"
	"time"
)

type Labyrinth struct {
	size_x uint
	size_y uint
	board  [][]bool
}

type Point struct {
	X, Y uint
}

func (l* Labyrinth) GetField(x, y uint) bool {
	return l.board[x][y]
}

func (l* Labyrinth) Dimensions() (x, y uint) {
	return l.size_x, l.size_y
}

func (l* Labyrinth) ForEach(function func(x, y uint, board *[][]bool)) {
	for j := uint(0); j < l.size_y; j++ {
		for i := uint(0); i < l.size_x; i++ {
			function(i, j, &l.board)
		}
	}
}


func Random(size_x, size_y uint) Labyrinth {
	rand.Seed(time.Now().Unix())
	board := newEmptyBoard(size_x, size_y)
	path := makePath(size_x, size_y)
	lab := Labyrinth{size_x: size_x, size_y: size_y, board: board}
	lab.ForEach(func(x, y uint, board *[][]bool) {
		(*board)[x][y] = !path[Point{x, y}]
	})
	return lab
}

func newEmptyBoard(sizeX, sizeY uint) [][]bool {
	board := make([][]bool, sizeX)
	for i := 0; uint(i) < sizeX; i++ {
		board[i] = make([]bool, sizeY)
	}
	return board
}

func makePath(sizeX, sizeY uint) map[Point]bool {
	curr := Point{0, 0}
	goal := Point{sizeX-1, sizeY-1}
	visited := map[Point]bool{
		curr: true,
		goal: true,
	}
	i := 0
	for !(curr.X == goal.X && curr.Y == goal.Y) {
		moveX, moveY := randMove(curr.X, curr.Y, sizeX-1, sizeY-1)
		//log.Printf("Move is (%v, %v)", moveX, moveY)
		curr = Point{
			uint(int(curr.X) + moveX),
			uint(int(curr.Y) + moveY),
		}
		visited[curr] = true
		//log.Printf("Iteration: %v, curr is %v", i, curr)
		i++
	}
	return visited
}

func randMove(currX, currY, maxX, maxY uint) (x, y int) {
	if rand.Intn(2) > 0 {
		return randAxisMove(currX, maxX), 0
	}
	return 0, randAxisMove(currY, maxY)
}

func randAxisMove(currAxisPosition, maxAxisPostion uint) int {
	move := rand.Intn(3) - 1
	castedAxis := int(currAxisPosition)
	castedMax := int(maxAxisPostion)
	for castedAxis + move < 0 || castedAxis + move > castedMax {
		move = rand.Intn(3) - 1
	}
	return move
}

