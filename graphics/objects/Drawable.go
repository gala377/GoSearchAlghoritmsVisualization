package objects

//
// Public interfaces
//

// Defines a type which renderer can draw
type Drawable interface {
	Draw()
}

// Defines a type which can be transformed
// in rendering space
type Transformable interface {
	Translate(x, y, z float32)
	Rotate(arc, x, y, z float32)
	Scale(x, y, z float32)

	GetPosition() (x, y, z float32)
	GetRotation() (x, y, z float32)
	GetScale() (x, y, z float32)
}