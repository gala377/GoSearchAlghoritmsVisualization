package MetaObjects

type Camera struct {

}

func (c *Camera) GetModel() []float32{
	panic("Implement me")
}

func (c *Camera) GetView() []float32 {
	panic("Implement me")
}

func (c *Camera) GetProjection() []float32 {
	panic("Implement me")
}