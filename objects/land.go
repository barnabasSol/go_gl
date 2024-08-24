package objects

import (
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Land struct {
	texture   helper.TextureId
	Vertices  []float32
	Positions []mgl32.Vec3
	VAO       helper.BufferId
}

func (land *Land) New() {
	land_file_path := filepath.Join("assets", "grass.jpg")
	land.texture = helper.LoadTextureAlphaJpeg(land_file_path)
	land.Vertices = []float32{
		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	land.Positions = make([]mgl32.Vec3, 0, 100*100)

	for x := -50; x <= 50; x++ {
		for z := -50; z <= 50; z++ {
			land.Positions = append(
				land.Positions,
				mgl32.Vec3{float32(x), -0.8, float32(z)},
			)
		}
	}
}

func (land *Land) LoadVertexAttribs() {
	land.VAO = helper.GenBindVertexArray(1)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 1)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, land.Vertices, gl.DYNAMIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()
}

func (land *Land) Renderer(shader_program *helper.Shader) {
	helper.BindVertextArray(land.VAO)
	helper.BindTexture(land.texture)
	for _, pos := range land.Positions {
		modelMatrix := mgl32.Ident4()
		modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
		shader_program.SetMat4("model", modelMatrix)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(land.Vertices)/5))
	}
}
