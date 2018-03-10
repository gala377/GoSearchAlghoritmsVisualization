package graphics


type Shader struct {}

func NewShader(frag, vert string) *Shader {
	panic("Implement me!")
}

func (sh *Shader) Use() {
	panic("Implement me!")
}

func (sh *Shader) GetProgramId () uint32  {
	panic("Implement me!")
}