package objects

import "github.com/go-gl/mathgl/mgl32"

type Land struct {
	Vertices  []float32
	Positions []mgl32.Vec3
}

func (land *Land) New() {
	land.Vertices = []float32{
		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}
	land.Positions = make([]mgl32.Vec3, 0, 20*20)

	for x := -10; x <= 10; x++ {
		for z := -10; z <= 10; z++ {
			land.Positions = append(land.Positions, mgl32.Vec3{float32(x), -0.8, float32(z)})
		}
	}

}
