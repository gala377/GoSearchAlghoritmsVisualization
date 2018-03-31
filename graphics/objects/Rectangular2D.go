package objects

import (
	"log"

	"github.com/gala377/SearchAlghorithms/graphics/objects/drawable"
)

type Rectangular struct {
	*drawable.Impl
	*RawObject
}

func NewSquare2D(size float32) *Rectangular {
	r := NewRawObject(
		[]float32{-size, size, 0.0, size, size, 0.0, size, -size, 0.0, -size, -size, 0.0},
		[]float32{0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0},
		[]uint32{0, 1, 3, 1, 2, 3},
	)
	r.CompileShaders("color/default.frag", "color/default.vert")
	return &Rectangular{
		Impl: &drawable.Impl{
			ConcreteDrawable: &drawableRawObjectImpl{
				rawModel: r,
			},
		},
		RawObject: r,
	}
}

func (sq *Rectangular) Draw() {
	log.Println("Called Rectangular2D draw")
	sq.Impl.Draw()
}

func (sq *Rectangular) SetShaderUniforms() {
	log.Println("Setting Rectangular2D shaderUniforms")
	sq.shader.Set4f("color", 0.0, 1.0, 0.0, 0.0)
}

