package segment

import (
	. "eagle/service/common"
	"github.com/gogf/gf/container/glist"
)

type ISegment interface {
	AtomSegment(charArray []Char, start int, end int) glist.List
	SimpleAtomSegment(charArray []Char, start int, end int) glist.List
	QuickAtomSegment(charArray []Char, start int, end int) glist.List // 快速原子分词
	Seg(text []Char) glist.List // 分词
	Seg2(text string) glist.List // 分词
	SegSentence(sentence []Char) glist.List
	Convert(vertexList glist.List, offsetEnabled bool) glist.List // 将一条路径转为最终结果
}


