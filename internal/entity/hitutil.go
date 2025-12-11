package entity

import (
	"math"

	"github.com/samber/lo"
)

func CircleHit(centerPos lo.Tuple2[float64, float64], radius float64, x float64, y float64) (bool, float64) {
	distance := math.Sqrt((x-centerPos.A)*(x-centerPos.A) + (y-centerPos.B)*(y-centerPos.B))
	return distance <= radius, distance
}

func EntityHit(e1 Shape, e2 Shape) bool {
	// e1 が e2 よりも左にある
	if e1.X+e1.Width < e2.X {
		return false
	}

	// e1 が e2 よりも右にある
	if e1.X > e2.X+e2.Width {
		return false
	}

	// e1 が e2 よりも上にある
	if e1.Y+e1.Height < e2.Y {
		return false
	}

	// e1 が e2 よりも下にある
	if e1.Y > e2.Y+e2.Height {
		return false
	}

	return true
}
