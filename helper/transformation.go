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

// func NewOrientation() *Orientation {

// 	orientation := &Orientation{
// 		X: randomFloat(rng),
// 		Y: randomFloat(rng),
// 		Z: randomFloat(rng),
// 	}

// 	// Randomly adjust one of the axes
// 	adjustAxis(rng, orientation)

// 	return orientation
// }

// func randomFloat(rng *rand.Rand) float32 {
// 	return rng.Float32() // Generates a float32 in the range [0.0, 1.0)
// }
