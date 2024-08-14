package helper

import (
	"github.com/veandco/go-sdl2/sdl"
)

var keyStates = sdl.GetKeyboardState()

func HandleInput(xv *float32, yv *float32) {
	if keyStates[sdl.SCANCODE_A] == 1 && keyStates[sdl.SCANCODE_W] == 1 {
		*xv += 0.0123
		*yv += 0.0123
	} else if keyStates[sdl.SCANCODE_D] == 1 && keyStates[sdl.SCANCODE_W] == 1 {
		*xv -= 0.0123
		*yv += 0.0123
	} else if keyStates[sdl.SCANCODE_A] == 1 && keyStates[sdl.SCANCODE_S] == 1 {
		*xv += 0.0123
		*yv -= 0.0123
	} else if keyStates[sdl.SCANCODE_D] == 1 && keyStates[sdl.SCANCODE_S] == 1 {
		*xv -= 0.0123
		*yv -= 0.0123
	} else if keyStates[sdl.SCANCODE_A] == 1 {
		*xv += 0.0123
	} else if keyStates[sdl.SCANCODE_D] == 1 {
		*xv -= 0.0123
	} else if keyStates[sdl.SCANCODE_W] == 1 {
		*yv += 0.0123
	} else if keyStates[sdl.SCANCODE_S] == 1 {
		*yv -= 0.0123
	}
}
