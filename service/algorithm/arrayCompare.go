package algorithm

import (
	"github.com/gogf/gf/util/gconv"
	"math"
)

type ArrayCompare struct {
}

// 比较数组A与B的大小关系
func (ac *ArrayCompare) Compare(arrayA []int64, arrayB []int64) int {
	len1 := len(arrayA)
	len2 := len(arrayB)
	lim := gconv.Int(math.Min(float64(len1), float64(len2)))

	k := 0
	for k < lim {
		var c1 int64 = arrayA[k]
		var c2 int64 = arrayB[k]
		if c1 != c2 {
			if c1 < c2 {
				return -1
			} else if c1 == c2 {
				return 0
			} else {
				return 1
			}
		}
		k++
	}

	return len1 - len2
}
