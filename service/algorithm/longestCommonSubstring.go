package algorithm

import . "gifs/service/common"

// 求最长公共字串的长度
// 最长公共子串（Longest Common Substring）指的是两个字符串中的最长公共子串，要求子串一定连续
type LongestCommonSubstring struct {

}

func (lcs *LongestCommonSubstring) Compute(str1 []Character, str2 []Character) int {
	size1 := len(str1)
	size2 := len(str2)
	if size1 == 0 || size2 == 0 {
		return 0
	}

	longest := 0

	for i := 0; i < size1; i++ {
		m := i
		n := 0
		length := 0
		for m < size1 && n < size2 {
			if str1[m] != str2[n] {
				length = 0
			}else{
				length++
				if longest < length {
					longest = length
				}
			}
			m++
			n++
		}
	}

	// 在str2中查找最长子串
	for j := 1; j < size2; j++ {
		m := 0
		n := j
		length := 0
		for m < size1 && n < size2 {
			if str1[m] != str2[n] {
				length = 0
			} else{
				length++
				if longest < length {
					longest = length
				}
			}
			m++
			n++
		}
	}

	return longest
}

func (lcs *LongestCommonSubstring) ComputeString(str1 String, str2 String) int {
	return lcs.Compute(str1.ToCharArray(), str2.ToCharArray())
}