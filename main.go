package main

import (
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 1280
const winHeight = 730

func main() {
	sdl_err := sdl.Init(sdl.INIT_EVERYTHING)
	if sdl_err != nil {
		panic(sdl_err)
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

	// fmt.Println(helper.GetVersion())

	vertexShaderPath := filepath.Join("shaders", "triangle.vert")
	fragmentShaderPath := filepath.Join("shaders", "triangle.frag")

	shader_program, err := helper.CreateProgram(vertexShaderPath, fragmentShaderPath)
	if err != nil {
		panic(err)
	}

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	helper.GenBindBuffer(gl.ARRAY_BUFFER, 1)
	VAO := helper.GenBindVertexArray(1)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	helper.UnbindVertexArray()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		helper.UseProgram(shader_program)
		helper.BindVertextArray(VAO)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.GLSwap()
		// helper.CheckForShaderChanges()
	}
}
