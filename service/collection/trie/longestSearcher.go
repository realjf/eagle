package trie

import . "eagle/service/common"

// 一个最长搜索工具（注意，当调用next()返回false后不应该继续调用next()，除非reset状态）
type LongestSearcher struct {
	Begin       int         // key的起点
	Length      int         // key的长度
	Index       int         // key的字典序坐标
	Value       interface{} // key对应的value
	charArray   []Char // 传入的字符数组
	i           int         // 上一个字符的下标
	arrayLength int         // charArray的长度，效率起见，开个变量

	check []int
	base  []int
	v     []interface{}
}

/**
 * 构造一个双数组搜索工具
 *
 * @param offset    搜索的起始位置
 * @param charArray 搜索的目标字符数组
 */
func NewLongestSearcher(offset int, charArray []Char) *LongestSearcher {
	return &LongestSearcher{
		charArray:   charArray,
		i:           offset,
		arrayLength: len(charArray),
		Begin:       offset,
	}
}

/**
 * 取出下一个命中输出
 *
 * @return 是否命中，当返回false表示搜索结束，否则使用公开的成员读取命中的详细信息
 */
func (ls *LongestSearcher) Next() bool {
	ls.Value = nil
	ls.Begin = ls.i
	var b int = ls.base[0]
	var n, p int

	for ; ; ls.i++ {
		if ls.i >= ls.arrayLength { // 指针到头了，将起点往前挪一个，重新开始，状态归零
			return ls.Value != nil
		}
		p = b + ls.charArray[ls.i].ToInt() + 1 // 状态转移 p = base[char[i-1]] + char[i] + 1
		if b == ls.check[p] {                  // base[char[i-1]] == check[base[char[i-1]] + char[i] + 1]
			b = ls.base[p] // 转移成功
		} else {
			if ls.Begin == ls.arrayLength {
				break
			}
			if ls.Value != nil {
				ls.i = ls.Begin + ls.Length // 输出最长词后，从该词语的下一个位置恢复扫描
				return true
			}

			ls.i = ls.Begin // 转移失败，也将起点往前挪一个，重新开始，状态归零
			ls.Begin++
			b = ls.base[0]
		}

		p = b
		n = ls.base[p]
		if b == ls.check[p] && n < 0 { // base[p] == check[p] && base[p] < 0 查到一个词
			ls.Length = ls.i - ls.Begin + 1
			ls.Index = -n - 1
			ls.Value = ls.v[ls.Index]
		}
	}

	return false
}
