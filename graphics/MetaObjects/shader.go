package MetaObjects

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Shader struct {
	frag uint32
	vert uint32

	program uint32
}


func NewShader(frag, vert string) (*Shader, error) {
	log.Println("Creating empty shader")
	sh := &Shader{}
	var err error
	log.Println("Compiling fragment shader...")
	if sh.frag, err = newConcreteShader("simple.frag", gl.FRAGMENT_SHADER); err != nil {
		log.Printf("Error (%v)", err)
		return sh, err
	}
	log.Println("Compiled.")
	log.Println("Compiling vertex shader...")
	if sh.vert, err = newConcreteShader("simple.vert", gl.VERTEX_SHADER); err != nil {
		log.Printf("Error (%v)", err)
		return sh, err
	}
	log.Println("Compiled.")
	log.Println("Creating shader program and linking shaders...")
	sh.program = gl.CreateProgram()
	gl.AttachShader(sh.program, sh.vert)
	gl.AttachShader(sh.program, sh.frag)
	gl.LinkProgram(sh.program)
	log.Println("Done.")

	var success int32
	if gl.GetProgramiv(sh.program, gl.LINK_STATUS, &success); !(success == 0) {
		infoLog := make([]byte, 4096)
		gl.GetProgramInfoLog(sh.program, int32(len(infoLog)), nil, &infoLog[0])
		log.Println("SHADER COMPILATION FAILED!")
		log.Printf("Direct couse:\n%v", string(infoLog))
		return sh, errors.New("shaders could not compile")
	}
	log.Println("Deleting shaders...")
	sh.delete()
	return sh, nil
}

func newConcreteShader(file string, shaderType uint32) (uint32, error) {
	sh := gl.CreateShader(shaderType)
	fileContent, err := readShaderFromFile(file)
	if err != nil {
		return 0, err
	}
	source, free := gl.Strs(fileContent)
	gl.ShaderSource(sh, 1, source, nil)
	free()
	gl.CompileShader(sh)
	return sh, nil
}

func readShaderFromFile(fileName string) (string, error) {
	absPath, err := filepath.Abs("graphics/shaders/")
	if err != nil {
		return "", err
	}
	src, err := ioutil.ReadFile(absPath + "/" + fileName)
	return string(src), err
}

func (sh* Shader) delete() {
	gl.DeleteShader(sh.frag)
	gl.DeleteShader(sh.vert)
}

func (sh *Shader) Use() {
	gl.UseProgram(sh.program)
}

func (sh *Shader) GetProgramId () uint32  {
	return sh.program
}


