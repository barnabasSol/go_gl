package helper

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Direction uint8

const (
	Forward Direction = iota
	Backward
	Left
	Right
	Nowhere
)

type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3

	WorldUp mgl32.Vec3

	Yaw              float32
	Pitch            float32
	MovementSpeed    float32
	MouseSensitivity float32
	Zoom             float32
}

func NewCamera(position mgl32.Vec3, world_up mgl32.Vec3, yaw, pitch, speed, sensitivity float32) *Camera {
	camera := &Camera{
		Position:         position,
		WorldUp:          world_up,
		Yaw:              yaw,
		Pitch:            pitch,
		MovementSpeed:    speed,
		MouseSensitivity: sensitivity,
	}
	camera.updateVectors()
	return camera
}

func (camera *Camera) UpdateCamera(
	direction Direction,
	deltaT float32,
	speed float32,
	x_offset, y_offset float32,
) {
	magnitude := camera.MovementSpeed * deltaT
	switch direction {
	case Forward:
		camera.Position = camera.Position.Add(camera.Front.Mul(magnitude))
	case Backward:
		camera.Position = camera.Position.Sub(camera.Front.Mul(magnitude))
	case Left:
		camera.Position = camera.Position.Sub(camera.Right.Mul(magnitude))
	case Right:
		camera.Position = camera.Position.Add(camera.Right.Mul(magnitude))
	case Nowhere:
	}
	x_offset *= camera.MouseSensitivity
	y_offset *= camera.MouseSensitivity

	camera.Yaw += x_offset
	camera.Pitch += y_offset

	camera.updateVectors()
}

func (camera *Camera) updateVectors() {
	front := mgl32.Vec3{float32(math.Cos(float64(mgl32.DegToRad(camera.Yaw) * float32(math.Cos(float64(mgl32.DegToRad(camera.Pitch))))))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Yaw)) * math.Cos(float64(mgl32.DegToRad(camera.Pitch))))),
	}

	camera.Front = front.Normalize()
	camera.Right = camera.Front.Cross(camera.WorldUp).Normalize()
	camera.Up = camera.Right.Cross(camera.Front).Normalize()
}

func (camera *Camera) GetViewMatrix() mgl32.Mat4 {
	center := camera.Position.Add(camera.Front)
	return mgl32.LookAt(
		camera.Position.X(),
		camera.Position.Y(),
		camera.Position.Z(),
		center.X(),
		center.Y(),
		center.Z(),
		camera.Up.X(),
		camera.Up.Y(),
		camera.Up.Z(),
	)

}
