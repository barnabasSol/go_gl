package helper

import (
	"math"
)

func Rotate(angle *float64, x *float32, y *float32) {
	cos := float32(math.Cos(*angle))
	sin := float32(math.Sin(*angle))

	*x = *x*cos - *y*sin
	*y = *x*sin + *y*cos
}

type Orientation struct {
	X float32
	Y float32
	Z float32
}
