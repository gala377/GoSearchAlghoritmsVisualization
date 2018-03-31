package objects

import (
	"log"

	"github.com/go-gl/gl/v4.3-core/gl"
)

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


// Drawable struct implementing ConcreteDrawable interface
// with objects.RawObject as its base
type drawableRawObjectImpl struct {
	rawModel *RawObject
}

func (impl *drawableRawObjectImpl) UseShader() {
	log.Println("Using rawObjImpl shader")
	impl.rawModel.shader.Use()
}

func (impl *drawableRawObjectImpl) SetTransformUniforms() {
	log.Println("Using rawObjImpl trans uniforms")
	impl.rawModel.setUniforms()
}

func (*drawableRawObjectImpl) SetShaderUniforms() {
	log.Println("Using rawObjImpl shaders uniforms")
}

func (impl *drawableRawObjectImpl) DrawElements() {
	log.Println("Using rawObjImpl draw elem")
	gl.BindVertexArray(impl.rawModel.VAO)
	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(impl.rawModel.indices)),
		gl.UNSIGNED_INT,
		nil)
	gl.BindVertexArray(0)
}