package graphics

import (
	"log"
	"runtime"

	"github.com/gala377/SearchAlghorithms/graphics/objects"
)
import "github.com/go-gl/glfw/v3.2/glfw"
import "github.com/go-gl/gl/v4.3-core/gl"

// TODO shared materials
// TODO Compound Drawable
// TODO Better performance


const WindowScalingFactor = 10.0

type Renderer struct {
	*Window
	objects []objects.Drawable

	polygonMode bool
	//TODO Change it later so there is a higher order struct
	//TODO Composing Renderer and inputManager, handling input
	input *InputManager
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
	r := &Renderer{
		&Window{},
		make([]objects.Drawable, 0),
		false,
		nil,
	}
	r.input = NewInputManager(r)
	return r, nil
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
	r.adjustObjectVertices(obj)
}

func (r* Renderer) adjustObjectVertices(obj objects.Drawable) {
	if casted, ok := obj.(objects.Transformable); ok {
		log.Println("Got transformable object... Scaling with window size...")
		casted.Scale(
			WindowScalingFactor/float32(r.width),
			WindowScalingFactor/float32(r.height),
			1.0 )
	}
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
	if r.input.GetKey(Key(glfw.KeyEscape)) == PRESSED {
		log.Println("ESCAPE pressed")
		r.window.SetShouldClose(true)
	}
	if r.input.GetKey(Key(glfw.KeyW)) == CLICKED {
		log.Println("W Clicked")
		if r.polygonMode {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		}
		r.polygonMode = !r.polygonMode
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
		log.Printf("Window size is (%d, %d) new width and height are (%d, %d)", r.width, r.height, width, height)
		for _, obj := range r.objects {
			r.rescaleObjectsVertices(obj, width, height)
		}
		r.width, r.height = uint32(width), uint32(height)
		gl.Viewport(0, 0, int32(width), int32(height))
	}
}

func (r *Renderer) rescaleObjectsVertices(obj objects.Drawable, width, height int) {
	if casted, ok := obj.(objects.Transformable); ok {
		log.Println("Got transformable object... Scaling with window size...")
		//x, y, z := casted.GetPosition()
		//casted.Translate(-x, -y, -z)
		casted.Scale(
			float32(r.width)/float32(width),
			float32(r.height)/float32(height),
			1.0 )
		//casted.Translate(x, y, z)
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


