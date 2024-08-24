package gamelogic

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/barnabasSol/go_gl/helper"
	"github.com/barnabasSol/go_gl/objects"
	"github.com/go-gl/mathgl/mgl32"
)

func GetHitEnemy(
	camera *helper.Camera,
	bim *objects.BulletInMotion,
	enemy_position *[]objects.ExtraEnemyProperty,
) (objects.ExtraEnemyProperty, error) {

	const gap = 1
	for _, prop := range *enemy_position {
		if bim.PosX >= prop.Position.X()-gap && bim.PosX <= prop.Position.X()+gap &&
			bim.PosY >= prop.Position.Y()-gap && bim.PosY <= prop.Position.Y()+gap &&
			bim.PosZ >= prop.Position.Z() && bim.PosZ <= prop.Position.Z()+gap {
			bim.PosZ = camera.Position.Z()
			return prop, nil
		}
	}
	return objects.ExtraEnemyProperty{}, errors.New("no enemy is hit nigger")
}

func MoveEnemies(enemies *[]objects.ExtraEnemyProperty, speed *float32) {
	var fucking_z float32 = -.025
	var fucking_y float32 = .025
	for i, prop := range *enemies {
		if !prop.IsHit {
			former_z_position := prop.Position.Z()
			former_z_position += *speed
			prop.Position = mgl32.Vec3{prop.Position.X(), prop.Position.Y(), former_z_position}
			(*enemies)[i] = prop
		} else {
			// println("whyy tho")
			former_z_position := prop.Position.Z()
			former_z_position += fucking_z
			former_y_position := prop.Position.Y()
			former_y_position += fucking_y
			prop.Position = mgl32.Vec3{prop.Position.X(), former_y_position, former_z_position}
			(*enemies)[i] = prop
		}
	}
}

func KillEnemy(
	hit_enemy *objects.ExtraEnemyProperty,
	enemies *[]objects.ExtraEnemyProperty,
	score *ScoreTrack,
) {
	score.KillCount += 1
	score.Points += 1
	for i, enemy := range *enemies {
		if enemy.Id == hit_enemy.Id {
			(*enemies)[i].IsHit = true
			break
		}
	}
}

func LevelUp(score *ScoreTrack, enemySpeed *float32, enemies *[]objects.ExtraEnemyProperty) {
	score.PassedMaxLevel += 1
	*enemySpeed *= 1.2
	spawnMore(enemies)
}

var init_count int = 4

func spawnMore(enemies *[]objects.ExtraEnemyProperty) {
	*enemies = (*enemies)[:0]

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	minX, maxX := float32(-12.0), float32(12.0)
	minZ, maxZ := float32(-30.0), float32(-10.0)
	y := float32(0.6)

	for i := 0; i < init_count*2; i++ {
		x := minX + rng.Float32()*(maxX-minX)
		z := minZ + rng.Float32()*(maxZ-minZ)

		id := i + 1
		(*enemies) = append(
			(*enemies),
			objects.ExtraEnemyProperty{
				Id:       uint32(id),
				IsHit:    false,
				Position: mgl32.Vec3{x, y, z},
			},
		)
	}

	init_count *= 2
}

func AllEnemiesAreHit(enemies *[]objects.ExtraEnemyProperty) bool {
	for _, enemy := range *enemies {
		if !enemy.IsHit || enemy.Position.Y() <= 5 {
			return false
		}
	}
	return true
}

func HandlePassedEnemies(enemies *[]objects.ExtraEnemyProperty, player_position *helper.Camera, scoreTrack *ScoreTrack) {
	for i, enemy := range *enemies {
		if enemy.Position.Z() > player_position.Position.Z() {
			scoreTrack.Points -= 1
			passedID := enemy.Id
			*enemies = append((*enemies)[:i], (*enemies)[i+1:]...)
			fmt.Printf("Enemy with ID %d passed you!\n", passedID)
		}
	}
}
