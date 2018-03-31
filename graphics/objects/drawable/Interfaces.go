package drawable

//
// Internal interfaces and structs
//

type ConcreteDrawable interface {
	UseShader()
	SetTransformUniforms()
	SetShaderUniforms()
	DrawElements()
}


