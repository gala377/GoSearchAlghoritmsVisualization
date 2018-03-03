package main

import (
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
	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)


	const window_width, window_height = 800, 600
	window, err := glfw.CreateWindow(window_width, window_height, "LearnOpenGL", nil, nil)
	defer glfw.Terminate()
	if err != nil {
		closer.Fatalln("Failed to create window:", err)
	}

	if err := gl.Init(); err != nil {
		closer.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(framebufferSizeCallback)


	gl.Viewport(0, 0, window_width, window_height)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	for !window.ShouldClose() {
		processInput(window)

		//rendering goes here
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()
		glfw.PollEvents()
	}


}

func framebufferSizeCallback(window *glfw.Window, width, height int) {
	gl.Viewport(0 ,0, int32(width), int32(height))
}


func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}