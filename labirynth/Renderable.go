package labyrinth

import (
	"log"

	"github.com/gala377/SearchAlghorithms/graphics/objects"
)
import glm "github.com/go-gl/mathgl/mgl32"

type Board struct {
	labyrinth Labyrinth

	squares []*objects.Rectangular
}

func NewBoard(l Labyrinth) Board {
	b := Board{
		labyrinth: l,
		squares: make([]*objects.Rectangular, 0),
	}
	for i, column := range l.board {
		for j := range column {
			cValue := float32((5*i*j) % 255)/255.0
			log.Printf("[%d: %d] Color is %v", j, i, cValue)
			b.squares = append(b.squares, objects.NewSquare2D(
				2.0, glm.Vec4{cValue, 0.2, 0.7, 0.0}))
			log.Printf("Position is %d, %d", i, j)
			b.squares[b.matCordsToLinear(uint32(j), uint32(i))].Translate(float32(i)/10, float32(j)/10, 0)
		}
	}
	return b
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
}

func (b *Board) Rotate(arc, x, y, z float32) {
	for _, square := range b.squares {
		square.Rotate(arc, x, y, z)
	}
}

func (b *Board) Scale(x, y, z float32) {
	for _, square := range b.squares {
		square.Scale(x, y, z)
	}
}

func (b *Board) GetPosition() (x, y, z float32) {
	panic("implement me")
}

func (b *Board) GetRotation() (x, y, z float32) {
	panic("implement me")
}

func (b *Board) GetScale() (x, y, z float32) {
	panic("implement me")
}

func (b *Board) matCordsToLinear(x, y uint32) uint32 {
	return x + y*uint32(b.labyrinth.size_y)
}



