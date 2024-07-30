package main

import (
	"fmt"

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
		0,
		100,
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

	vertex_shader := helper.CreateShader(vertex_shader_src, gl.VERTEX_SHADER)
	fragment_shader := helper.CreateShader(fragment_shader_src, gl.VERTEX_SHADER)

	shader_program := helper.CreateProgram(vertex_shader, fragment_shader)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	VBO := helper.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := helper.GenBindVertexArray()
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(0)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(uint32(shader_program))
		gl.BindVertexArray(uint32(VAO))
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.GLSwap()
	}
}
