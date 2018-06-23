package labyrinth

import (
	"github.com/gala377/SearchAlghorithms/graphics/objects"
)
import glm "github.com/go-gl/mathgl/mgl32"

type Board struct {
	labyrinth Labyrinth

	squares []*objects.Rectangular

	xPos, yPos float32
}

type Settings struct {
	WallColor, PathColor glm.Vec4
	SquareSize float32
	Shift float32
}

func DefaultSettings() Settings {
	return Settings{
		glm.Vec4{0.5, 0.5, 0.5, 0.0},
		glm.Vec4{1.0, 1.0, 1.0, 0.0},
		2.0,
		10,
	}
}

func NewBoard(l Labyrinth, settings Settings) *Board {
	b := &Board{
		labyrinth: l,
		squares: make([]*objects.Rectangular, 0),
		xPos: 0,
		yPos: 0,
	}
	b.labyrinth.ForEach(func(x, y uint, board *[][]bool) {
		cValue := settings.PathColor
		if (*board)[x][y] {
			cValue = settings.WallColor
		}
		b.squares = append(b.squares, objects.NewSquare2D(settings.SquareSize, cValue))
		b.squares[b.matCordsToLinear(x, y)].Translate(
			float32(x)/settings.Shift, float32(y)/settings.Shift, 0)
	})
	return b
}


func (board* Board) SetColor(x, y uint, r, g, b float32) {
	board.squares[board.matCordsToLinear(x, y)].Color = glm.Vec4{r, g, b, 0.0}
}


func (b *Board) Draw() {
	for _, square := range b.squares {
		square.Draw()
	}
}


func (b *Board) Translate(x, y, z float32) {
	for _, square := range b.squares {
		square.Translate(x, y, z)
	}
	b.xPos += x
	b.yPos += y
}

func (b *Board) Rotate(arc, x, y, z float32) {
	for _, square := range b.squares {
		square.Rotate(arc, x, y, z)
	}
}

func (b *Board) Scale(x, y, z float32) {
	for _, square := range b.squares {
		//xPos, yPos, zPos := b.GetPosition()
		//square.Translate(-xPos, -yPos, -zPos)
		square.Scale(x, y, z)
		//square.Translate(xPos, yPos, zPos)
	}
}

func (b *Board) GetPosition() (x, y, z float32) {
	return b.xPos, b.yPos, 0.0
}

func (b *Board) GetRotation() (x, y, z float32) {
	panic("implement me")
}

func (b *Board) GetScale() (x, y, z float32) {
	panic("implement me")
}

func (b *Board) SetPosition(x, y, z float32) {
	panic("implement me")
}

func (b *Board) SetRotation(arc, x, y, z float32) {
	panic("implement me")
}

func (b *Board) SetScale(x, y, z float32) {
	panic("implement me")
}

func (b *Board) matCordsToLinear(x, y uint) uint {
	return x + y*b.labyrinth.size_y
}



