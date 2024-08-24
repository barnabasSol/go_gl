package main

import (
	"path/filepath"
	"time"

	"github.com/barnabasSol/go_gl/gamelogic"
	"github.com/barnabasSol/go_gl/helper"
	"github.com/barnabasSol/go_gl/objects"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 1280
const winHeight = 730

var keyStates = sdl.GetKeyboardState()

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	defer sdl.Quit()
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)

	scoreTrack := gamelogic.NewScore(0, 0)

	window, wind_err := sdl.CreateWindow(
		"Hello GoGl",
		50,
		30,
		winWidth,
		winHeight,
		sdl.WINDOW_OPENGL)

	if wind_err != nil {
		panic(wind_err)
	}

	window.GLCreateContext()
	defer window.Destroy()

	gl.Init()
	gl.Enable(gl.DEPTH_TEST)

	vertexShaderPath := filepath.Join("shaders", "first.vert")
	fragmentShaderPath := filepath.Join("shaders", "quad_tex.frag")

	shader_program, err := helper.NewShader(vertexShaderPath, fragmentShaderPath)
	if err != nil {
		panic(err)
	}

	world_up := mgl32.Vec3{0.0, 1.0, 0.0}
	position := mgl32.Vec3{0.0, 0.0, 3.0}
	camera := helper.NewCamera(position, world_up, -90, .0, 0.01, .4)

	var land objects.Land
	var enemy objects.Enemey
	var player objects.Player
	var bullet objects.Bullet

	land.New()
	enemy.New()
	player.New()
	bullet.New()

	land.LoadVertexAttribs()
	enemy.LoadVertexAttribs()
	player.LoadVertexAttribs()
	bullet.LoadVertexAttribs()

	var elapsedTime float32
	var enemySpeed float32 = 0.01111
	bim := objects.BulletInMotion{
		PosX: camera.Position.X(),
		PosY: camera.Position.Y(),
		PosZ: camera.Position.Z(),
	}
	prevMouseX, prevMouseY, _ := sdl.GetMouseState()
	for {
		frameStart := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				gamelogic.PrintGameStats(scoreTrack)
				return
			}
		}

		if keyStates[sdl.SCANCODE_SPACE] != 0 {
			bullet.IsFired = true
		} else {
			bullet.IsFired = false
		}

		mouseX, mouseY, _ := sdl.GetMouseState()

		var direction helper.Direction = helper.Nowhere

		if keyStates[sdl.SCANCODE_A] != 0 {
			bim.PosX = camera.Position.X()
			direction = helper.Left
		}
		if keyStates[sdl.SCANCODE_D] != 0 {
			bim.PosX = camera.Position.X()
			direction = helper.Right
		}
		if keyStates[sdl.SCANCODE_W] != 0 {
			bim.PosZ = camera.Position.Z()
			direction = helper.Forward
		}
		if keyStates[sdl.SCANCODE_S] != 0 {
			bim.PosZ = camera.Position.Z()
			direction = helper.Backward
		}

		if camera.Position.Y() <= 0.3 {
			var newY float32 = 0.4
			camera.Position = mgl32.Vec3{camera.Position.X(), newY, camera.Position.Z()}
		}

		camera.UpdateCamera(
			direction,
			elapsedTime,
			camera.MovementSpeed,
			float32(mouseX-prevMouseX),
			-float32(mouseY-prevMouseY),
		)

		prevMouseX = mouseX
		prevMouseY = mouseY

		gl.ClearColor(0.53, 0.81, 0.92, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		shader_program.Use()

		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 200.0)
		viewMatrix := camera.GetViewMatrix()
		shader_program.SetMat4("projection", projectionMatrix)
		shader_program.SetMat4("view", viewMatrix)

		land.Renderer(shader_program)

		enemy.Renderer(shader_program)

		gamelogic.MoveEnemies(&enemy.Extras, &enemySpeed)
		if gamelogic.AllEnemiesAreHit(&enemy.Extras) || len(enemy.Extras) < 1 {
			println("omgg all enemies are hit")
			gamelogic.LevelUp(scoreTrack, &enemySpeed, &enemy.Extras)
		}

		player.Renderer(camera, shader_program)

		bullet.Renderer(camera, shader_program, &bim)

		hit_enemy, err := gamelogic.GetHitEnemy(camera, &bim, &enemy.Extras)
		if err == nil {
			gamelogic.KillEnemy(&hit_enemy, &enemy.Extras, scoreTrack)
		}
		gamelogic.HandlePassedEnemies(&enemy.Extras, camera, scoreTrack)
		window.GLSwap()
		shader_program.CheckForShaderChanges()
		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
	}
}
