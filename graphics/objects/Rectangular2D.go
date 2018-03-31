package objects

import glm "github.com/go-gl/mathgl/mgl32"

type Rectangular struct {
	*ExtendedRawObject

	Color glm.Vec4
}

func NewSquare2D(size float32, color glm.Vec4) *Rectangular {
	e := NewExtendedRawObject(
		[]float32{-size, size, 0.0, size, size, 0.0, size, -size, 0.0, -size, -size, 0.0},
		[]float32{0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0},
		[]uint32{0, 1, 3, 1, 2, 3},
	)
	rect := &Rectangular{
		ExtendedRawObject: e,
		Color: color,
	}

	e.setShaderUniforms = func (r *RawObject) {
		r.shader.SetV4f("color", rect.Color)
	}
	e.CompileShaders("color/default.frag", "color/default.vert")
	return rect
}
