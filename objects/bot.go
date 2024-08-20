package objects

import (
	"github.com/barnabasSol/go_gl/helper"
	"github.com/go-gl/mathgl/mgl32"
)

type Bot struct {
	Vertices  []float32
	Positions []mgl32.Vec3
	VAO       helper.BufferId
}

func (bot *Bot) New() {

}
