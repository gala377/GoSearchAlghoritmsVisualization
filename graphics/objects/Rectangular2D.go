package objects

func NewSquare2D(size float32) *RawObject {
	r := NewRawObject(
		[]float32{-size, size, 0.0, size, size, 0.0, size, -size, 0.0, -size, -size, 0.0},
		[]float32{0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0},
		[]uint32{0, 1, 3, 1, 2, 3},
	)
	r.CompileShaders("simple.frag", "simple.vert")
	return r
}
