package algorithm

import (
	"github.com/gogf/gf/util/gconv"
	"math"
)

type ArrayDistance struct {
}

// 求两个集合中最相近的两个数
func (ad *ArrayDistance) ComputeMinimumDistance(arrayA []int64, arrayB []int64) int64 {
	aIndex := 0
	bIndex := 0
	var min int64 = gconv.Int64(math.Abs(gconv.Float64(arrayA[0] - arrayB[0])))
	for {
		if arrayA[aIndex] > arrayB[bIndex] {
			bIndex++
		} else {
			aIndex++
		}
		if aIndex >= len(arrayA) || bIndex >= len(arrayB) {
			break
		}
		min2 := gconv.Int64(math.Abs(gconv.Float64(arrayA[aIndex] - arrayB[bIndex])))
		if min2 < min {
			min = min2
		}
	}

	return min
}

func (ad *ArrayDistance) ComputeAverageDistance(arrayA []int64, arrayB []int64) int64 {
	var totalA int64 = 0
	var totalB int64 = 0
	for _, a := range arrayA {
		totalA += a
	}
	for _, b := range arrayB {
		totalB += b
	}

	return gconv.Int64(
		math.Abs(
			gconv.Float64(
				totalA/gconv.Int64(len(arrayA)) - totalB/gconv.Int64(len(arrayB)))))
}
