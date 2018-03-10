package objects


type Drawable interface {
	Draw()
}

type Transformable interface {
	Translate(x, y, z float32)
	Rotate(arc, x, y, z float32)
	Scale(x, y, z float32)

	GetPosition() (x, y, z float32)
	GetRotation() (x, y, z float32)
	GetScale() (x, y, z float32)
}