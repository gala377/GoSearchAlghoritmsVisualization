package main

import (
	"io/ioutil"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/xlab/closer"
)


func init() {
	runtime.LockOSThread()
}

func main() {
	window := initGLFW()
	defer glfw.Terminate()

	prog := initOpenGL()

	for !window.ShouldClose() {
		processInput(window)

		//rendering goes here
		draw(window, prog)
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func initGLFW() *glfw.Window {
	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	const window_width, window_height = 800, 600
	window, err := glfw.CreateWindow(window_width, window_height, "LearnOpenGL", nil, nil)
	if err != nil {
		closer.Fatalln("Failed to create window:", err)
	}

	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(
		func(window *glfw.Window, width, height int) {
			gl.Viewport(0, 0, int32(width), int32(height))
		})

	gl.Viewport(0, 0, window_width, window_height)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		closer.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	window.SwapBuffers()
	glfw.PollEvents()
}


type Triangle struct {
	VBO uint32
	VAO uint32
	vertices []float32

}

func NewTriangle() *Triangle {
	t := Triangle{
		vertices: []float32{
			-0.5, -0.5, 0.0,
			0.5, -0.5, 0.0,
			0.0,  0.5, 0.0,
		},
	}
	gl.GenBuffers(1, &t.VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, t.VBO)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		4*len(t.vertices),
		gl.Ptr(t.vertices),
		gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &t.VAO)
	gl.BindVertexArray(t.VAO)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, t.VBO)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return &t
}

func loadVertexShader() []byte {
	file, err := ioutil.ReadFile("trianngle.vert")
	if err != nil {
		closer.Fatalln(err)
	}
	return file
}