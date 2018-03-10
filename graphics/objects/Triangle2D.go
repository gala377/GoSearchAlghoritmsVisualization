package objects

func NewTriangle2D() *RawObject {
	return NewRawObject(
		[]float32{-0.5, -0.5, 0.0, 0.5, -0.5, 0.0, 0.0, 0.5, 0.0},
		[]float32{0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0},
		[]uint32{0, 1, 2},
	)
}