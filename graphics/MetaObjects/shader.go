package MetaObjects

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	//"log"
	"path/filepath"
	"strings"

	"github.com/go-gl/gl/v4.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
)

// TODO caching shader uniform locations
// TODO chaching compiled shaders - ShaderManager
type Shader struct {
	frag uint32
	vert uint32

	program uint32

	cachedUniformLocations map[string]int32
}

// Move to shader manager in later release
var shaders = make(map[string]*Shader)

// Use this instead of NewShader for material caching
func GetShader(path string) (*Shader, error) {
	sh, ok := shaders[path]
	if !ok {
		sh, err := NewShader(path+".frag", path+".vert")
		if err != nil {
			return nil, err
		}
		shaders[path] = sh
	}
	return sh, nil
}

func NewShader(frag, vert string) (*Shader, error) {
	log.Println("Creating empty shader")
	sh := &Shader{}
	var err error
	log.Println("Compiling fragment shader...")
	if sh.frag, err = compileShader(frag, gl.FRAGMENT_SHADER); err != nil {
		log.Printf("Error (%v)", err)
		return sh, err
	}
	log.Println("Compiled.")
	log.Println("Compiling vertex shader...")
	if sh.vert, err = compileShader(vert, gl.VERTEX_SHADER); err != nil {
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
	log.Println("Checking for errors")
	err = sh.checkForLinkingAndCompileErrors()
	log.Println("Deleting shaders...")
	sh.delete()
	return sh, err
}

func compileShader(file string, shaderType uint32) (uint32, error) {
	sh := gl.CreateShader(shaderType)
	fileContent, err := readShaderFromFile(file)
	if err != nil {
		return 0, err
	}
	source, free := gl.Strs(fileContent)
	gl.ShaderSource(sh, 1, source, nil)
	free()
	gl.CompileShader(sh)

	var success int32
	if gl.GetShaderiv(sh, gl.COMPILE_STATUS, &success); success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(sh, gl.INFO_LOG_LENGTH, &logLength)

		logInfo := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(sh, logLength, nil, gl.Str(logInfo))

		return 0, fmt.Errorf("Failed to compile %v: %v", source, logInfo)
	}

	return sh, nil
}

func readShaderFromFile(fileName string) (string, error) {
	absPath, err := filepath.Abs("graphics/shaders/")
	if err != nil {
		return "", err
	}
	src, err := ioutil.ReadFile(absPath + "/" + fileName)
	src = append(src, 0)
	return string(src), err
}

func (sh *Shader) checkForLinkingAndCompileErrors() error {
	var success int32
	if gl.GetProgramiv(sh.program, gl.LINK_STATUS, &success); success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(sh.program, gl.INFO_LOG_LENGTH, &logLength)
		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(sh.program, logLength, nil, gl.Str(infoLog))
		log.Println("SHADER LINKING FAILED!")
		log.Printf("Direct couse:\n%v", string(infoLog))
		return errors.New("shaders could not link")
	}
	if gl.GetProgramiv(sh.program, gl.COMPILE_STATUS, &success); success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(sh.program, gl.INFO_LOG_LENGTH, &logLength)
		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(sh.program, logLength, nil, gl.Str(infoLog))
		log.Println("SHADER LINKING FAILED!")
		log.Printf("Direct couse:\n%v", string(infoLog))
		return errors.New("shaders could not link")
	}
	return nil
}

func (sh *Shader) delete() {
	gl.DeleteShader(sh.frag)
	gl.DeleteShader(sh.vert)
}

//
// Public Interface
//

func (sh *Shader) Use() {
	gl.UseProgram(sh.program)
}

func (sh *Shader) GetProgramId() uint32 {
	return sh.program
}

func (sh *Shader) SetBool(name string, value bool) {
	var boolValue int32
	if value {
		boolValue = 1
	} else {
		boolValue = 0
	}
	gl.Uniform1i(sh.getUniformLocation(name), boolValue)
}

func (sh *Shader) SetInt(name string, value int) {
	gl.Uniform1i(sh.getUniformLocation(name), int32(value))
}

func (sh *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(sh.program, gl.Str(name)), value)
}

func (sh *Shader) Set4f(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(sh.getUniformLocation(name), v0, v1, v2, v3)
}

func (sh *Shader) SetV4f(name string, vec glm.Vec4) {
	v0, v1, v2, v3 := vec.Elem()
	sh.Set4f(name, v0, v1, v2, v3)
}

func (sh *Shader) getUniformLocation(name string) int32 {
	loc, ok := sh.cachedUniformLocations[name]
	if !ok {
		name += "\x00"
		sh.cachedUniformLocations[name] = gl.GetUniformLocation(sh.program, gl.Str(name))
		loc = sh.cachedUniformLocations[name]
	}
	return loc
}
