package gamelogic

import (
	"errors"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/barnabasSol/go_gl/objects"
	"github.com/go-gl/mathgl/mgl32"
)

func HitEnemyIndex(
	camera *helper.Camera,
	bim *objects.BulletInMotion,
	enemy_position *[]mgl32.Vec3,
) (int16, error) {

	const gap = 1
	for i, enemy := range *enemy_position {
		if bim.PosX >= enemy.X()-gap && bim.PosX <= enemy.X()+gap &&
			bim.PosY >= enemy.Y()-gap && bim.PosY <= enemy.Y()+gap &&
			bim.PosZ >= enemy.Z() && bim.PosZ <= enemy.Z()+gap {
			bim.PosZ = camera.Position.Z()
			return int16(i), nil
		}
	}
	return -1, errors.New("no enemy is hit nigger")
}

func KillEnemy(
	enemy_index int16,
	enemies *[]mgl32.Vec3,
) {
	println("im hittttttttttt")
	var magnitude float32 = 0.2111
	former_position := (*enemies)[enemy_index].X()
	former_position -= magnitude
	(*enemies)[enemy_index] = mgl32.Vec3{former_position, (*enemies)[enemy_index].Y(), (*enemies)[enemy_index].Z()}
}

func CheckIfAllEnemiesAreOut(enemies *[]mgl32.Vec3) bool {
	return len(*enemies) < 1
}

func SpawnMore() {}
