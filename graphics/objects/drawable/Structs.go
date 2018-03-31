package drawable

// Drawable struct implementing ConcreteDrawable interface
// in Draw function
type Impl struct {
	ConcreteDrawable
}

func (impl *Impl) Draw() {
	//log.Println("Using drawable.Impl Draw()")
	impl.UseShader()
	impl.SetTransformUniforms()
	impl.SetShaderUniforms()
	impl.DrawElements()
}
