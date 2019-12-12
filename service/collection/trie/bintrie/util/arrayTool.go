package util

import (
	"eagle/service/collection/trie/bintrie"
	. "eagle/service/common"
)

/**
 * 二分查找
 * @param branches 数组
 * @param node 要查找的node
 * @return 数组下标，小于0表示没找到
 */
func BinarySearch(branches []*bintrie.BaseNode, node bintrie.BaseNode) int {
	high := len(branches) -1
	if len(branches) < 1 {
		return high
	}
	low := 0
	for low <= high {
		var mid int = int(uint(low + high) >> 1)
		var cmp int = int(branches[mid].CompareTo2(node))

		if cmp < 0 {
			low = mid + 1
		}else if cmp > 0 {
			high = mid - 1
		}else {
			return mid
		}
	}
	return -(low + 1)
}

func BinarySearch2(branches []*bintrie.BaseNode, node Char) int {
	high := len(branches) - 1
	if len(branches) < 1 {
		return high
	}
	low := 0
	for low < high {
		var mid int = int(uint(low + high) >> 1)
		var cmp int = int(branches[mid].CompareTo(node))

		if cmp < 0 {
			low = mid + 1
		}else if cmp > 0 {
			high = mid -1
		}else {
			return mid
		}
	}
	return -(low + 1)
}
