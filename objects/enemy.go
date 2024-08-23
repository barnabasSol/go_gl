package objects

import (
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type PoppedEnemies []mgl32.Vec3

type ExtrProperty struct {
	Id       uint32
	Position mgl32.Vec3
	IsHit    bool
}

type Enemey struct {
	ModelMatrix   mgl32.Mat4
	textures      []helper.TextureId
	Vertices      []float32
	Extra         []ExtrProperty
	VAO           helper.BufferId
	MovementSpeed float32
}

func (enemy *Enemey) New() {
	green_file_path := filepath.Join("assets", "green_monster.jpg")
	green_texture := helper.LoadTextureAlphaJpeg(green_file_path)
	white_file_path := filepath.Join("assets", "white_monster.jpg")
	white_texture := helper.LoadTextureAlphaJpeg(white_file_path)

	enemy.textures = []helper.TextureId{green_texture, white_texture}

	enemy.Vertices = []float32{
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
	enemy.Extra = []ExtrProperty{
		{Id: 1, IsHit: false, Position: mgl32.Vec3{-4.0, 0.6, -7.0}},
		{Id: 2, IsHit: false, Position: mgl32.Vec3{1.0, 0.6, -7.0}},
		{Id: 3, IsHit: false, Position: mgl32.Vec3{4.0, 0.6, -7.0}},
		{Id: 4, IsHit: false, Position: mgl32.Vec3{8.0, 0.6, -7.0}},
	}

	// enemy.Positions = []mgl32.Vec3{
	// 	{-4.0, 0.6, -7.0},
	// 	{1.0, 0.6, -7.0},
	// 	{4.0, 0.6, -7.0},
	// 	{8.0, 0.6, -7.0},
	// 	{-4.0, 0.6, -10.0},
	// 	{1.0, 0.6, -10.0},
	// 	{4.0, 0.6, -10.0},
	// 	{8.0, 0.6, -10.0},
	// }

}

func (enemy *Enemey) LoadVertexAttribs() {
	enemy.VAO = helper.GenBindVertexArray(2)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 2)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, enemy.Vertices, gl.DYNAMIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()
}

var increment float32 = 0.0444
var angle float32 = 0

func (enemy *Enemey) Renderer(shader_program *helper.Shader) {
	helper.BindVertextArray(enemy.VAO)
	for i, prop := range enemy.Extra {
		if i%2 == 0 {
			helper.BindTexture(enemy.textures[0])
		} else {
			helper.BindTexture(enemy.textures[1])
		}
		angle += increment
		enemy.ModelMatrix = mgl32.Ident4()
		enemy.ModelMatrix = mgl32.HomogRotate3DX(mgl32.DegToRad(angle)).Mul4(enemy.ModelMatrix)
		enemy.ModelMatrix = mgl32.Translate3D(prop.Position.X(), prop.Position.Y(), prop.Position.Z()).Mul4(enemy.ModelMatrix)
		shader_program.SetMat4("model", enemy.ModelMatrix)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}
}
