package utils

import (
	"math"

	"github.com/samber/lo"
)

func CircleHit(centerPos lo.Tuple2[float64, float64], radius float64, x float64, y float64) (bool, float64) {
	distance := math.Sqrt((x-centerPos.A)*(x-centerPos.A) + (y-centerPos.B)*(y-centerPos.B))
	return distance <= radius, distance
}
