package objects

import (
	"log"

	//"log"
	"github.com/gala377/SearchAlghorithms/graphics/MetaObjects"
	"github.com/go-gl/gl/v4.3-core/gl"
)
import glm "github.com/go-gl/mathgl/mgl32"

const FLOAT_SIZE = 4
const INT_SIZE = 4

type RawObject struct {
	VBO uint32
	EBO uint32
	VAO uint32

	trans    glm.Mat4
	position glm.Vec3
	rotation glm.Vec3
	scale    glm.Vec3

	shader *MetaObjects.Shader

	vertices []float32
	indices  []uint32

	camera *MetaObjects.Camera
}

//
// Creation
//

func NewRawObject(verts, normals []float32, indics []uint32) *RawObject {
	r := emptyRawObject()
	r.indices = make([]uint32, len(indics))
	copy(r.indices, indics)
	log.Printf("Idicies copied: %v", r.indices)
	for i := 0; i < (len(verts) / 3); i++ {
		curr := i * 3
		r.vertices = append(r.vertices, verts[curr:curr+3]...)
		r.vertices = append(r.vertices, normals[curr:curr+3]...)
	}
	r.bindBuffers()
	return r
}

func emptyRawObject() *RawObject {
	r := &RawObject{
		vertices: make([]float32, 0),
		indices:  make([]uint32, 0),
		trans:    glm.Ident4(),
		position: glm.Vec3{0.0, 0.0, 0.0},
		rotation: glm.Vec3{0.0, 0.0, 0.0},
		scale:    glm.Vec3{1.0, 1.0, 1.0},
	}
	gl.GenBuffers(1, &r.VBO)
	gl.GenBuffers(1, &r.EBO)
	gl.GenVertexArrays(1, &r.VAO)
	return r
}

func (r *RawObject) bindBuffers() {
	gl.BindVertexArray(r.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.VBO)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(r.vertices)*FLOAT_SIZE,
		gl.Ptr(r.vertices),
		gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.EBO)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(r.indices)*INT_SIZE,
		gl.Ptr(r.indices),
		gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*FLOAT_SIZE, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*FLOAT_SIZE, gl.Ptr(r.vertices[3:]))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

//
// RawObject interface
//


func (r *RawObject) SetShader(path string) (err error) {
	r.shader, err = MetaObjects.GetShader(path)
	return err
}

// TODO change to SetShade
func (r *RawObject) CompileShaders(frag, vert string) (err error) {
	r.shader, err = MetaObjects.NewShader(frag, vert)
	return err
}

//
// Drawable impl
//

func (r *RawObject) Draw() {
	r.shader.Use()

	r.setUniforms()

	gl.BindVertexArray(r.VAO)

	gl.DrawElements(gl.TRIANGLES, int32(len(r.indices)), gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (r *RawObject) setUniforms() {
	transformLoc := gl.GetUniformLocation(r.shader.GetProgramId(), gl.Str("transform\x00"))
	//modelLoc := gl.GetUniformLocation(r.shader.GetProgramId(), gl.Str("model\x00"))
	//viewLoc := gl.GetUniformLocation(r.shader.GetProgramId(), gl.Str("view\x00"))
	//projectionLoc := gl.GetUniformLocation(r.shader.GetProgramId(), gl.Str("projection\x00"))

	gl.UniformMatrix4fv(transformLoc, 1, false, &r.trans[0])
	//gl.UniformMatrix4fv(modelLoc, 1, false, &r.camera.GetModel()[0])
	//gl.UniformMatrix4fv(viewLoc, 1, false, &r.camera.GetView()[0])
	//gl.UniformMatrix4fv(projectionLoc, 1, false, &r.camera.GetProjection()[0])
}

//
// Transformable impl
//

func (r *RawObject) Translate(x, y, z float32) {
	r.trans = r.trans.Add(glm.Translate3D(x, y, z))
	r.position = r.position.Add(glm.Vec3{x, y, z})
}

func (r *RawObject) Rotate(arc, x, y, z float32) {
	r.trans = r.trans.Mul4(glm.HomogRotate3D(arc, glm.Vec3{x, y, z}))
	r.rotation = r.rotation.Add(glm.Vec3{x, y, z}.Mul(arc))
}

func (r *RawObject) Scale(x, y, z float32) {
	r.trans = r.trans.Mul4(glm.Scale3D(x, y, z))
	r.scale = glm.Vec3{r.scale.X() * x, r.scale.Y() * y, r.scale.Z() * z}
}

func (r *RawObject) GetPosition() (x, y, z float32) {
	return r.position.Elem()
}

func (r *RawObject) GetRotation() (x, y, z float32) {
	return r.rotation.Elem()
}

func (r *RawObject) GetScale() (x, y, z float32) {
	return r.scale.Elem()
}
