package helper

// func HandleInput(xv *float32, yv *float32) {
// 	if keyStates[sdl.SCANCODE_A] == 1 && keyStates[sdl.SCANCODE_W] == 1 {
// 		*xv += 0.0123
// 		*yv += 0.0123
// 	} else if keyStates[sdl.SCANCODE_D] == 1 && keyStates[sdl.SCANCODE_W] == 1 {
// 		*xv -= 0.0123
// 		*yv += 0.0123
// 	} else if keyStates[sdl.SCANCODE_A] == 1 && keyStates[sdl.SCANCODE_S] == 1 {
// 		*xv += 0.0123
// 		*yv -= 0.0123
// 	} else if keyStates[sdl.SCANCODE_D] == 1 && keyStates[sdl.SCANCODE_S] == 1 {
// 		*xv -= 0.0123
// 		*yv -= 0.0123
// 	} else if keyStates[sdl.SCANCODE_A] == 1 {
// 		*xv += 0.0123
// 	} else if keyStates[sdl.SCANCODE_D] == 1 {
// 		*xv -= 0.0123
// 	} else if keyStates[sdl.SCANCODE_W] == 1 {
// 		*yv += 0.0123
// 	} else if keyStates[sdl.SCANCODE_S] == 1 {
// 		*yv -= 0.0123
// 	}
// }

// func HandleCameraInput(camera *Camera, x_offset, y_offset float32) {
// 	frameStart := time.Now()
// 	var direction Direction = Nowhere
// 	if keyStates[sdl.SCANCODE_A] != 0 {
// 		direction = Left
// 	}
// 	if keyStates[sdl.SCANCODE_D] != 0 {
// 		direction = Right
// 	}
// 	if keyStates[sdl.SCANCODE_W] != 0 {
// 		direction = Forward
// 	}
// 	if keyStates[sdl.SCANCODE_S] != 0 {
// 		direction = Backward
// 	}
// 	elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
// 	fmt.Println(elapsedTime)
// 	camera.UpdateCamera(direction, elapsedTime, camera.MovementSpeed, x_offset, y_offset)
// }
