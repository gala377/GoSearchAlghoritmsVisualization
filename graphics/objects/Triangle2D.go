package objects

func NewTriangle2D() (*RawObject) {
	r := NewRawObject(
		[]float32{-0.5, -0.5, 0.0, 0.5, -0.5, 0.0, 0.0, 0.5, 0.0},
		[]float32{0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0},
		[]uint32{0, 1, 2},
	)
	err := r.CompileShaders("simple.frag", "simple.vert")
	if err != nil {
		panic(err)
	}
	return r
}