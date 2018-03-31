package objects

import "github.com/go-gl/gl/v4.3-core/gl"


type ExtendedRawObject struct {
	*RawObject
	// Assign your own functions to control drawing process
	useShaderProgram func(object *RawObject)
	setTransformUniforms func(object *RawObject)
	setShaderUniforms func(object *RawObject)
	drawElements func(object *RawObject)
}

//
// Drawable implementation
//

func (e* ExtendedRawObject) Draw() {
	e.useShaderProgram(e.RawObject)
	e.setTransformUniforms(e.RawObject)
	e.setShaderUniforms(e.RawObject)
	e.drawElements(e.RawObject)
}

//
// Creation
//

func NewExtendedRawObject(verts, normals []float32, indics []uint32) *ExtendedRawObject {
	e := ExtendedRawObject{
		RawObject: NewRawObject(verts, normals, indics),
	}
	e.assignDefaultFunctions()
	return &e
}

func (e *ExtendedRawObject) assignDefaultFunctions() {

	e.useShaderProgram = func(r *RawObject) {
		r.shader.Use()
	}
	e.setTransformUniforms = func(r *RawObject) {
		r.setUniforms()
	}
	e.setShaderUniforms = func(r *RawObject) {}
	e.drawElements = func(r *RawObject) {
		gl.BindVertexArray(r.VAO)

		gl.DrawElements(gl.TRIANGLES, int32(len(r.indices)), gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)
	}
}