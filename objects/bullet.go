package objects

import "github.com/go-gl/mathgl/mgl32"

type Bullet struct {
	Vertices  []float32
	Positions []mgl32.Vec3
}

func (bullet *Bullet) New() {
	// Bullet Vertices
	// Bullet Vertices
	bullet.Vertices = []float32{
		// Front face
		-0.05, 0.2, 0.05, 0.0, 1.0, // 0. Top left
		0.05, 0.2, 0.05, 1.0, 1.0, // 1. Top right
		0.05, -0.2, 0.05, 1.0, 0.0, // 2. Bottom right
		-0.05, -0.2, 0.05, 0.0, 0.0, // 3. Bottom left

		// Back face
		-0.05, 0.2, -0.05, 0.0, 1.0, // 4. Top left
		0.05, 0.2, -0.05, 1.0, 1.0, // 5. Top right
		0.05, -0.2, -0.05, 1.0, 0.0, // 6. Bottom right
		-0.05, -0.2, -0.05, 0.0, 0.0, // 7. Bottom left

		// Top face
		-0.05, 0.2, 0.05, 0.0, 0.0, // 8. Top left
		0.05, 0.2, 0.05, 1.0, 0.0, // 9. Top right
		0.05, 0.2, -0.05, 1.0, 1.0, // 10. Top right
		-0.05, 0.2, -0.05, 0.0, 1.0, // 11. Top left

		// Bottom face
		-0.05, -0.2, 0.05, 0.0, 0.0, // 12. Bottom left
		0.05, -0.2, 0.05, 1.0, 0.0, // 13. Bottom right
		0.05, -0.2, -0.05, 1.0, 1.0, // 14. Bottom right
		-0.05, -0.2, -0.05, 0.0, 1.0, // 15. Bottom left
	}

	bullet.Positions = []mgl32.Vec3{
		{0.0, 0.0, -7.0},
	}

}
