package gamelogic

import (
	"errors"
	"fmt"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/barnabasSol/go_gl/objects"
	"github.com/go-gl/mathgl/mgl32"
)

func HitEnemyId(
	camera *helper.Camera,
	bim *objects.BulletInMotion,
	enemy_position *[]objects.ExtrProperty,
) (objects.ExtrProperty, error) {

	const gap = 1
	for _, prop := range *enemy_position {
		if bim.PosX >= prop.Position.X()-gap && bim.PosX <= prop.Position.X()+gap &&
			bim.PosY >= prop.Position.Y()-gap && bim.PosY <= prop.Position.Y()+gap &&
			bim.PosZ >= prop.Position.Z() && bim.PosZ <= prop.Position.Z()+gap {
			bim.PosZ = camera.Position.Z()
			return prop, nil
		}
	}
	return objects.ExtrProperty{}, errors.New("no enemy is hit nigger")
}

func MoveEnemies(enemies *[]objects.ExtrProperty, speed float32) {
	var fucking_z float32 = -.025
	var fucking_y float32 = .025
	for i, prop := range *enemies {
		if !prop.IsHit {
			former_z_position := prop.Position.Z()
			former_z_position += speed
			prop.Position = mgl32.Vec3{prop.Position.X(), prop.Position.Y(), former_z_position}
			(*enemies)[i] = prop
		} else {
			println("whyy tho")
			former_z_position := prop.Position.Z()
			former_z_position += fucking_z
			former_y_position := prop.Position.Y()
			former_y_position += fucking_y
			prop.Position = mgl32.Vec3{prop.Position.X(), former_y_position, former_z_position}
			(*enemies)[i] = prop
		}
	}

}

// func defuseEnemy(popped_enemies *[]objects.ExtrProperty, speed float32) {
// 	speed *= -1
// 	for i, enemy := range *popped_enemies {
// 		if enemy.IsHit {
// 			former_x_position := enemy.Position.X()
// 			former_x_position += speed
// 			enemy.Position = mgl32.Vec3{former_x_position, enemy.Position.Y(), enemy.Position.Z()}
// 			(*popped_enemies)[i].Position = enemy.Position
// 			println("omggggggggg")
// 		}
// 	}
// }

func KillEnemy(
	hit_enemy *objects.ExtrProperty,
	enemies *[]objects.ExtrProperty,
) {
	fmt.Println(hit_enemy.IsHit)
	hit_enemy.IsHit = true

	println("im hittttttttttt", hit_enemy.IsHit)

	for i, enemy := range *enemies {
		if enemy.Id == hit_enemy.Id {
			(*enemies)[i].IsHit = true
			break
		}
	}

	fmt.Println(hit_enemy.IsHit)
	// defuseEnemy(enemies, 0.00111)
}

func SpawnMore() {}

func AllEnemiesPassedPlayer() {

}
