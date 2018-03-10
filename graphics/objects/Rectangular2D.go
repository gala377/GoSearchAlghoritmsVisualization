package objects

func NewSquare2D(size float32) *RawObject {
	return NewRawObject(
		[]float32{-size, size, 0.0, size, size, 0.0, -size, -size, 0.0, size, -size, 0.0},
		[]float32{0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0},
		[]uint32{0, 1, 3, 1, 2, 3},
	)
}
