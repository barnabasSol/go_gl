package main

import (
	"fmt"
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
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
		sdl.WINDOW_OPENGL,
	)

	if wind_err != nil {
		panic(wind_err)
	}
	window.GLCreateContext()
	defer window.Destroy()
	gl.Init()

	fmt.Println(helper.GetVersion())

	vertexShaderPath := filepath.Join("shaders", "first.vert")
	fragmentShaderPath := filepath.Join("shaders", "quad_tex.frag")

	shader_program, err := helper.NewShader(vertexShaderPath, fragmentShaderPath)
	if err != nil {
		panic(err)
	}
	tex_file_path := filepath.Join("assets", "gold.png")
	texture := helper.LoadTextureAlpha(tex_file_path)

	vertices := []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	cubePositions := []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -10.0},
		{2.0, 5.0, -14.0},
	}

	helper.GenBindBuffer(gl.ARRAY_BUFFER, 1)
	VAO := helper.GenBindVertexArray(1)
	// helper.GenBindBuffer(gl.ELEMENT_ARRAY_BUFFER, 1)
	// helper.BufferDataInt(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

	helper.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()

	//for vertex shader

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		shader_program.Use()

		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(65.0), float32(winWidth)/float32(winHeight), 0.1, 100.0)
		// viewMatrix := mgl32.Ident4()
		viewMatrix := mgl32.Translate3D(0.0, 0.0, -3.0)
		shader_program.SetMat4("projection", projectionMatrix)

		shader_program.SetMat4("view", viewMatrix)

		helper.BindTexture(texture)
		helper.BindVertextArray(VAO)
		for _, pos := range cubePositions {
			modelMatrix := mgl32.Ident4()
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			// var angle float32 = 20.0 * float32(i)
			// modelMatrix = mgl32.HomogRotate3DY(mgl32.DegToRad(angle))
			shader_program.SetMat4("model", modelMatrix)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		window.GLSwap()
		shader_program.CheckForShaderChanges()
	}
}
