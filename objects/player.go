package objects

import (
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Player struct {
	texture   helper.TextureId
	Vertices  []float32
	Positions []mgl32.Vec3
	VAO       helper.BufferId
}

func (player *Player) New() {

	var gold_file_path string = filepath.Join("assets", "gold.png")
	player.texture = helper.LoadTextureAlphaPng(gold_file_path)

	player.Vertices = []float32{
		// Cylinder sides
		// Bottom face
		-0.5, -0.5, -0.5, 0.0, 1.0, // Bottom left back
		0.5, -0.5, -0.5, 1.0, 1.0, // Bottom right back
		0.5, -0.5, 0.5, 1.0, 0.0, // Bottom right front
		0.5, -0.5, 0.5, 1.0, 0.0, // Bottom right front
		-0.5, -0.5, 0.5, 0.0, 0.0, // Bottom left front
		-0.5, -0.5, -0.5, 0.0, 1.0, // Bottom left back

		// Top face
		-0.5, 0.5, -0.5, 0.0, 1.0, // Top left back
		0.5, 0.5, -0.5, 1.0, 1.0, // Top right back
		0.5, 0.5, 0.5, 1.0, 0.0, // Top right front
		0.5, 0.5, 0.5, 1.0, 0.0, // Top right front
		-0.5, 0.5, 0.5, 0.0, 0.0, // Top left front
		-0.5, 0.5, -0.5, 0.0, 1.0, // Top left back

		// Cylinder sides
		// Front face (with texture coordinates)
		-0.5, 0.5, 0.5, 0.0, 0.0, // Top left front
		0.5, 0.5, 0.5, 1.0, 0.0, // Top right front
		0.5, -0.5, 0.5, 1.0, 1.0, // Bottom right front
		0.5, -0.5, 0.5, 1.0, 1.0, // Bottom right front
		-0.5, -0.5, 0.5, 0.0, 1.0, // Bottom left front
		-0.5, 0.5, 0.5, 0.0, 0.0, // Top left front

		// Back face
		-0.5, 0.5, -0.5, 0.0, 0.0, // Top left back
		0.5, 0.5, -0.5, 1.0, 0.0, // Top right back
		0.5, -0.5, -0.5, 1.0, 1.0, // Bottom right back
		0.5, -0.5, -0.5, 1.0, 1.0, // Bottom right back
		-0.5, -0.5, -0.5, 0.0, 1.0, // Bottom left back
		-0.5, 0.5, -0.5, 0.0, 0.0, // Top left back

		// Left face
		-0.5, 0.5, -0.5, 0.0, 0.0, // Top left back
		-0.5, 0.5, 0.5, 1.0, 0.0, // Top left front
		-0.5, -0.5, 0.5, 1.0, 1.0, // Bottom left front
		-0.5, -0.5, 0.5, 1.0, 1.0, // Bottom left front
		-0.5, -0.5, -0.5, 0.0, 1.0, // Bottom left back
		-0.5, 0.5, -0.5, 0.0, 0.0, // Top left back

		// Right face
		0.5, 0.5, -0.5, 0.0, 0.0, // Top right back
		0.5, 0.5, 0.5, 1.0, 0.0, // Top right front
		0.5, -0.5, 0.5, 1.0, 1.0, // Bottom right front
		0.5, -0.5, 0.5, 1.0, 1.0, // Bottom right front
		0.5, -0.5, -0.5, 0.0, 1.0, // Bottom right back
		0.5, 0.5, -0.5, 0.0, 0.0, // Top right back
	}

	player.Positions = []mgl32.Vec3{
		{0.0, 0.7, -7.0},
	}

}

func (player *Player) LoadVertexAttribs() {
	player.VAO = helper.GenBindVertexArray(3)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 3)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, player.Vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()
}

func (player *Player) Renderer(camera *helper.Camera, shader_program *helper.Shader) {
	helper.BindVertextArray(player.VAO)
	helper.BindTexture(player.texture)
	modelMatrix := mgl32.Ident4()
	modelMatrix = mgl32.Translate3D(camera.Position.X(), camera.Position.Y()-.8, camera.Position.Z()).Mul4(modelMatrix)
	shader_program.SetMat4("model", modelMatrix)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
}
