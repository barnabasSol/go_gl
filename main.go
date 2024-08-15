package main

import (
	"fmt"
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/barnabasSol/go_gl/objects"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 1280
const winHeight = 730

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	defer sdl.Quit()
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)

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

	fmt.Println(helper.GetVersion())

	vertexShaderPath := filepath.Join("shaders", "first.vert")
	fragmentShaderPath := filepath.Join("shaders", "quad_tex.frag")

	shader_program, err := helper.NewShader(vertexShaderPath, fragmentShaderPath)
	if err != nil {
		panic(err)
	}
	grass_file_path := filepath.Join("assets", "land.jpeg")
	grass_texture := helper.LoadTextureAlphaJpeg(grass_file_path)
	gold_file_path := filepath.Join("assets", "gold.png")
	gold_texture := helper.LoadTextureAlphaPng(gold_file_path)

	var land objects.Land
	var cube objects.Cube

	land.New()
	cube.New()

	// Land
	landVAO := helper.GenBindVertexArray(1)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 1)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, land.Vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()

	// Cube VAO and VBO
	cubeVAO := helper.GenBindVertexArray(2)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 2)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, cube.Vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()

	var (
		x_pos float32 = 0
		z_pos float32 = 0
	)
	cubeX := cube.Positions[0].X()
	cubeZ := cube.Positions[1].Z()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		shader_program.Use()

		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 200.0)
		helper.HandleInput(&x_pos, &z_pos)
		viewMatrix := mgl32.Translate3D(x_pos, 0.0, z_pos)
		shader_program.SetMat4("projection", projectionMatrix)
		shader_program.SetMat4("view", viewMatrix)

		helper.BindVertextArray(landVAO)
		helper.BindTexture(grass_texture)
		for _, pos := range land.Positions {
			modelMatrix := mgl32.Ident4()
			// var angle float32 = 60.0 * float32(i)
			// modelMatrix = mgl32.HomogRotate3D(mgl32.DegToRad(angle), mgl32.Vec3{1.0, 0.3, 0.5}).Mul4(modelMatrix)
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			shader_program.SetMat4("model", modelMatrix)
			gl.DrawArrays(gl.TRIANGLES, 0, int32(len(land.Vertices)/5))
		}

		helper.BindTexture(gold_texture)
		helper.BindVertextArray(cubeVAO)
		for _, pos := range cube.Positions {
			modelMatrix := mgl32.Ident4()
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			shader_program.SetMat4("model", modelMatrix)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		cubeX += .003
		if cubeX > 4.0 {
			cubeX = -4.0
		}
		cubeZ += .006
		if cubeZ > 4.0 {
			cubeZ = -7.0
		}
		cube.Positions[0] = mgl32.Vec3{cubeX, cube.Positions[0].Y(), cube.Positions[0].Z()}
		cube.Positions[1] = mgl32.Vec3{cube.Positions[1].X(), cube.Positions[1].Y(), cubeZ}
		window.GLSwap()
		shader_program.CheckForShaderChanges()
	}
}
