package labyrinth

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
	return Labyrinth{}
}
