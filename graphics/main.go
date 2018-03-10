package graphics

import (
	"runtime"
)
import "github.com/go-gl/glfw/v3.2/glfw"
import "github.com/go-gl/gl/v4.3-core/gl"

type Renderer struct {
	*Window
}

type Window struct {
	window *glfw.Window
	width  uint32
	height uint32
	title  string
}

//
// Initialization
//

func New() (*Renderer, error) {
	runtime.LockOSThread()
	if err := initGLFW(); err != nil {
		return nil, err
	}
	return &Renderer{
		&Window{},
	}, nil
}

func initGLFW() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	return nil
}

//
//	Window interface
//

func (w *Window) SetWindowSize(width, height uint32) {
	w.height = height
	w.width = width
}

func (w *Window) SetWindowTitle(title string) {
	w.title = title
}

func (w *Window) SetWindow(width, height uint32, title string) {
	w.SetWindowSize(width, height)
	w.SetWindowTitle(title)
}

func (w *Window) GetWindow() (*glfw.Window, error) {
	if w.window != nil {
		return w.window, nil
	}
	var err error
	w.window, err = glfw.CreateWindow(int(w.width), int(w.height), w.title, nil, nil)
	if err != nil {
		return nil, err
	}
	w.window.MakeContextCurrent()
	err = initGL(w.width, w.height)
	return w.window, err

}

func initGL(width, height uint32) error {
	if err := gl.Init(); err != nil {
		return err
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	return nil
}

