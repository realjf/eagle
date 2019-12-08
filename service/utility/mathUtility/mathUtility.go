package mathUtility

import . "eagle/service/segment/common"

func Sum(vals ...int) int {
	sum := 0
	for _, x := range vals {
		sum += x
	}
	return sum
}

/**
 * 从一个词到另一个词的词的花费
 *
 * @param from 前面的词
 * @param to   后面的词
 * @return 分数
 */
func CalculateWeight(from Vertex, to Vertex) float64 {
	frequency := from.GetAttribute().TotalFrequency
	if frequency == 0 {
		frequency = 1 // 防止发生除零错误
	}
	nTwoWordsFreq :=
}


