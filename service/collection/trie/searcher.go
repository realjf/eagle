package trie

import . "eagle/service/common"

// 一个搜索工具（注意，当调用next()返回false后不应该继续调用next()，除非reset状态）
type Searcher struct {
	Begin int // key的起点
	Length int // key的长度
	Index int // key的字典序坐标
	Value interface{} // key对应的value
	charArray []Char // 传入的字符数组
	last int // 上一个node位置
	i int // 上一个字符的下标
	arrayLength int // charArray的长度，效率起见，开个变量

	LongestSearcher
}

func NewSearcher(offset int, charArray []Char) *Searcher {
	searcher := &Searcher{
		charArray: charArray,
		i: offset,
		arrayLength:len(charArray),
	}

	searcher.last = searcher.base[0]
	// A trick，如果文本长度为0的话，调用next()时，会带来越界的问题。
	// 所以我要在第一次调用next()的时候触发begin == arrayLength进而返回false。
	// 当然也可以改成begin >= arrayLength，不过我觉得操作符>=的效率低于==
	if searcher.arrayLength == 0 {
		searcher.Begin = -1
	}else{
		searcher.Begin = offset
	}

	return searcher
}

/**
 * 取出下一个命中输出
 *
 * @return 是否命中，当返回false表示搜索结束，否则使用公开的成员读取命中的详细信息
 */
func (s *Searcher) Next() bool {
	b := s.last
	var n, p int

	for ; ; s.i++ {
		if s.i == s.arrayLength {// 指针到头了，将起点往前挪一个，重新开始，状态归零
			s.Begin++
			if s.Begin == s.arrayLength {
				break
			}
			s.i = s.Begin
			b = s.base[0]
		}

		p = b + s.charArray[s.i].ToInt() + 1 // 状态转移 p = base[char[i-1]] + char[i] + 1
		if b == s.check[p] {// base[char[i-1]] == check[base[char[i-1]] + char[i] + 1]
			b = s.base[p]// 转移成功
		}else{
			s.i = s.Begin// 转移失败，也将起点往前挪一个，重新开始，状态归零
			s.Begin++
			if s.Begin == s.arrayLength {
				break
			}
			b = s.base[0]
			continue
		}
		p = b
		n = s.base[p]
		if b == s.check[p] && n < 0 { // base[p] == check[p] && base[p] < 0 查到一个词
			s.Length = s.i - s.Begin + 1
			s.Index = -n -1
			s.Value = s.v[s.Index]
			s.last = b
			s.i++
			return true
		}
	}

	return false
}

