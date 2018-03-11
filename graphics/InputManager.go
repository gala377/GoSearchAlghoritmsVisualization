package graphics

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyState int
type Key glfw.Key

const (
	CLICKED  KeyState = 0
	PRESSED  KeyState = 1
	RELEASED KeyState = 2
	FREE     KeyState = 3
)

type InputManager struct {
	renderer *Renderer

	buttons map[Key]KeyState
}

func NewInputManager(renderer *Renderer) *InputManager {
	return &InputManager{
		renderer: renderer,
		buttons: make(map[Key]KeyState),
	}
}

func (im *InputManager) GetKey(key Key) KeyState {
	previous := im.getKeyState(key)
	isPressed, err := im.isGLFWKeyPressed(key)
	if err != nil {
		// TODO change it to logging in later realease
		panic(err)
	}
	im.progressState(key, previous, isPressed)
	return im.buttons[key]
}

func (im *InputManager) getKeyState(key Key) KeyState {
	state, ok := im.buttons[key]
	//When it's first time accessing the key
	if !ok {
		im.buttons[key] = FREE
		state = im.buttons[key]
	}
	return state
}

func (im *InputManager) isGLFWKeyPressed(key Key) (bool, error) {
	win, err := im.renderer.GetWindow()
	if err != nil {
		return false, err
	}
	return win.GetKey(glfw.Key(key)) == glfw.Press, nil
}

func (im *InputManager) progressState(key Key, previous KeyState, isClicked bool) {
	switch previous {
	case CLICKED:
		if isClicked {
			im.buttons[key] = PRESSED
		} else {
			im.buttons[key] = RELEASED
		}
	case PRESSED:
		if !isClicked {
			im.buttons[key] = RELEASED
		}
	case RELEASED:
		if isClicked {
			im.buttons[key] = PRESSED
		} else {
			im.buttons[key] = FREE
		}
	case FREE:
		if isClicked {
			im.buttons[key] = CLICKED
		}
	}
}
