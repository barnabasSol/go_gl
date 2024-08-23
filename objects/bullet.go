package objects

import (
	"path/filepath"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Bullet struct {
	texture   helper.TextureId
	IsFired   bool
	ShotSpeed float32
	Vertices  []float32
	Position  mgl32.Vec3
	VAO       helper.BufferId
}

type BulletInMotion struct {
	PosX float32
	PosY float32
	PosZ float32
}

func (bullet *Bullet) New() {

	gold_file_path := filepath.Join("assets", "gold.png")
	bullet.texture = helper.LoadTextureAlphaPng(gold_file_path)

	bullet.IsFired = false
	bullet.ShotSpeed = .2222
	bullet.Vertices = []float32{
		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

}

func (bullet *Bullet) LoadVertexAttribs() {
	bullet.VAO = helper.GenBindVertexArray(3)
	helper.GenBindBuffer(gl.ARRAY_BUFFER, 1)
	helper.BufferDataFloat(gl.ARRAY_BUFFER, bullet.Vertices, gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)
	helper.UnbindVertexArray()
}

func (bullet *Bullet) Renderer(camera *helper.Camera, shader_program *helper.Shader, bim *BulletInMotion) {
	helper.BindVertextArray(bullet.VAO)
	helper.BindTexture(bullet.texture)

	if bullet.IsFired {
		bim.PosZ -= float32(bullet.ShotSpeed)
		var firing_range float32 = -20
		if bim.PosZ <= firing_range {
			bim.PosZ = camera.Position.Z()
			bullet.IsFired = false
		}
		modelMatrix := mgl32.Ident4()
		modelMatrix = mgl32.Translate3D(bim.PosX, camera.Position.Y(), bim.PosZ).Mul4(modelMatrix)
		shader_program.SetMat4("model", modelMatrix)
	} else {
		bim.PosZ = camera.Position.Z()
		modelMatrix := mgl32.Ident4()
		modelMatrix = mgl32.Translate3D(camera.Position.X(), camera.Position.Y()-.4, camera.Position.Z()).Mul4(modelMatrix)
		shader_program.SetMat4("model", modelMatrix)
	}
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(bullet.Vertices)/5))
}
