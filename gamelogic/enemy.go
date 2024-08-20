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
) (uint8, error) {
	const space = 1
	for i, enemy := range *enemy_position {
		if bim.PosX >= enemy.X()-space && bim.PosX <= enemy.X()+space &&
			bim.PosY >= enemy.Y()-space && bim.PosY <= enemy.Y()+space &&
			bim.PosZ >= enemy.Z() && bim.PosZ <= enemy.Z()+space {
			bim.PosZ = camera.Position.Z()
			return uint8(i), nil
		}
	}
	return 0, errors.New("no enemy is hit nigger")
}

func KillEnemy(
	enemy_index uint8,
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

func SpawnMore() {

}
