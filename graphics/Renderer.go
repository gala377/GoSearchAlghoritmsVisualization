package graphics

import (
	"fmt"
	"log"
	"runtime"

	"github.com/gala377/SearchAlghorithms/graphics/objects"
)
import "github.com/go-gl/glfw/v3.2/glfw"
import "github.com/go-gl/gl/v4.3-core/gl"

type Renderer struct {
	*Window
	objects []objects.Drawable
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
	log.Println("Locking thread...")
	runtime.LockOSThread()
	log.Println("Locked")
	log.Println("Initializing GLFW...")
	if err := initGLFW(); err != nil {
		log.Fatalf("Error: %v", err )
		return nil, err
	}
	log.Println("Creating empty renderer...")
	return &Renderer{
		&Window{},
		make([]objects.Drawable, 0),
	}, nil
}

func initGLFW() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	log.Println("Setting Window hints...")
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	log.Println("Window hints set.")
	return nil
}

//
// Renderer Interface
//

func (r *Renderer) AddObject(obj objects.Drawable) {
	r.objects = append(r.objects, obj)
}

func (r *Renderer) Terminate() {
	log.Println("Terminating GLFW")
	glfw.Terminate()
}

//I think creating a window should connect an appropriate callback
//So this function will be trashed in later refactoring
func (r *Renderer) ConnectCallbacks() {
	r.window.SetFramebufferSizeCallback(r.frameBufferSizeCallback())
}

func (r *Renderer) Draw() {
	r.processInput()

	r.render()

	r.window.SwapBuffers()
	glfw.PollEvents()
}

func (r *Renderer) processInput() {
	if r.window.GetKey(glfw.KeyEscape) == glfw.Press {
		r.window.SetShouldClose(true)
	}
}

func (r* Renderer) render() {
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range r.objects {
		obj.Draw()
	}
}

func (r *Renderer) frameBufferSizeCallback() glfw.FramebufferSizeCallback {
	return func (window *glfw.Window, width, height int) {
		r.SetWindowTitle(fmt.Sprintf("(%v, %v)", width, height))
		gl.Viewport(0, 0, int32(width), int32(height))
	}
}


//
//	Window interface
//

func (w *Window) SetWindowSize(width, height uint32) {
	w.height = height
	w.width = width
	log.Printf("Window Size set (%v, %v)", width, height)
}

func (w *Window) SetWindowTitle(title string) {
	w.title = title
	if w.window != nil {
		w.window.SetTitle(w.title)
	}
	log.Printf("Window Title set (%v)", title)
}

func (w *Window) SetWindow(width, height uint32, title string) {
	w.SetWindowSize(width, height)
	w.SetWindowTitle(title)
}

func (w *Window) GetWindow() (*glfw.Window, error) {
	log.Println("Returning window")
	if w.window != nil {
		return w.window, nil
	}
	log.Println("Window doesn't exists!")
	log.Println("Creating new glfw window...")
	var err error
	w.window, err = glfw.CreateWindow(int(w.width), int(w.height), w.title, nil, nil)
	if err != nil {
		log.Printf("Error (%v)", err)
		return nil, err
	}
	log.Println("Window created. Setting as current context...")
	w.window.MakeContextCurrent()
	log.Println("Set.")
	log.Println("Initializing GL...")
	err = initGL(w.width, w.height)
	log.Println("Initialized.")
	return w.window, err

}

func initGL(width, height uint32) error {
	log.Println("Calling gl.Init()...")
	if err := gl.Init(); err != nil {
		log.Printf("Error (%v)", err)
		return err
	}
	log.Println("Initialized!")
	log.Println("Setting viewport...")
	gl.Viewport(0, 0, int32(width), int32(height))
	log.Println("Set.")
	return nil
}


