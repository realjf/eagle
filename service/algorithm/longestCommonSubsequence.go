package algorithm

import (
	. "gifs/service/common"
	"github.com/gogf/gf/util/gconv"
	"math"
)

// 最长公共子序列（Longest Common Subsequence）指的是两个字符串中的最长公共子序列，不要求子序列连续。
type LongestCommonSubsequence struct {
}

func (lcs *LongestCommonSubsequence) Compute(str1 []Character, str2 []Character) int {
	substringLength1 := len(str1)
	substringLength2 := len(str2)

	// 构造二维数组记录子问题A[i]和B[j]的LCS的长度
	var opt [][]int = make([][]int, substringLength1+1)
	for i := 0; i < substringLength1+1; i++ {
		opt[i] = make([]int, substringLength2+1)
	}

	// 从后向前，动态规划计算所有子问题。也可从前到后。
	for i := substringLength1 - 1; i >= 0; i-- {
		for j := substringLength2 - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				opt[i][j] = opt[i+1][j+1] + 1 // 状态转移方程
			} else {
				// 状态转移方程
				opt[i][j] = gconv.Int(math.Max(gconv.Float64(opt[i+1][j]), gconv.Float64(opt[i][j+1])))
			}
		}
	}

	return opt[0][0]
}

func (lcs *LongestCommonSubsequence) ComputeString(str1 String, str2 String) int {
	return lcs.Compute(str1.ToCharArray(), str2.ToCharArray())
}
