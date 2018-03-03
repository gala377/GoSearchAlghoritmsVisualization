package main

import (
	"runtime"

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
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)


	const window_width, window_height = 800, 600
	window, err := glfw.CreateWindow(window_width, window_height, "LearnOpenGL", nil, nil)
	if err != nil {
		closer.Fatalln("Failed to create window:", err)
	}

	window.MakeContextCurrent()
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glfw.Terminate()
}
