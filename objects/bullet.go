package objects

import (
	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Bullet struct {
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
