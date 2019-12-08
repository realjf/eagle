package coreBiGramTableDictionary

import "eagle/service/dictionary/coreDictionary"

var (
	GCoreBiGramTableDictionary *CoreBiGramTableDictionary
)

func init() {
	if GCoreBiGramTableDictionary == nil {
		GCoreBiGramTableDictionary = NewCoreBiGramTableDictionary()
	}
}


// 核心词典的二元接续词典，采用整型储存，高性能
type CoreBiGramTableDictionary struct {
	Start []int
	Pair []int
}

func NewCoreBiGramTableDictionary() *CoreBiGramTableDictionary {
	return &CoreBiGramTableDictionary{}
}

func (c *CoreBiGramTableDictionary) BinarySearch(a []int, fromIndex int, length int, key int) int {
	low := fromIndex
	high := fromIndex + length - 1

	for low <= high {
		var mid int = int(uint(low + high) >> 1)
		midVal := a[mid << 1]

		if midVal < key {
			low = mid + 1
		}else if midVal > key {
			high = mid - 1
		}else{
			return mid
		}
	}
	return -(low + 1)
}

/**
 * 获取共现频次
 * @param idA 第一个词的id
 * @param idB 第二个词的id
 * @return 共现频次
 */
func (c *CoreBiGramTableDictionary) GetBiFrequency(a string, b string) int {
	idA := coreDictionary.GCoreDictionary.Trie.ExactMatchSearch(a)
	if idA == -1 {
		return 0
	}
	idB := coreDictionary.GCoreDictionary.Trie.ExactMatchSearch(b)
	if idB == -1 {
		return 0
	}
	index := c.BinarySearch(c.Pair, c.Start[idA], c.Start[idA+1] - c.Start[idA], idB)
	if index < 0 {
		return 0
	}
	index = index << 1
	return c.Pair[index+1]
}

/**
 * 获取共现频次
 * @param idA 第一个词的id
 * @param idB 第二个词的id
 * @return 共现频次
 */
func (c *CoreBiGramTableDictionary) GetBiFrequency2(idA int, idB int) int {
	if idA < 0 {
		return -idA
	}
	if idB < 0 {
		return -idB
	}
	index := c.BinarySearch(c.Pair, c.Start[idA], c.Start[idA+1] - c.Start[idA], idB)
	if index < 0 {
		return 0
	}
	index = index << 1
	return c.Pair[index+1]
}

