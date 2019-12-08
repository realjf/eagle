package segment

import (
	. "eagle/service/common"
	. "eagle/service/segment/common"
	"github.com/gogf/gf/container/glist"
)

type Segment struct {

}

func (s *Segment) SimpleAtomSegment(charArray []Char, start int, end int) glist.List {
	return *glist.New(true)
}

func (s *Segment) QuickAtomSegment(charArray []Char, start int, end int) glist.List {
	return *glist.New(true)
}

func (s *Segment) Seg(text []Char) glist.List {
	// 正规化字符

	return *glist.New(true)
}

func (s *Segment) SegSentence(sentence []Char) glist.List {
	return *glist.New(true)
}

// 将一条路径转为最终结果
func (s *Segment) Convert(vertexList glist.List, offsetEnabled bool) glist.List {
	length := vertexList.Size() - 2
	resultList := glist.New(true)
	if offsetEnabled {
		offset := 0
		for i := 0; i <length; i++ {
			resultList.Iterator(func(e *glist.Element) bool {
				var vertex Vertex = e.Value.(Vertex)
				var term Term = s.ConvertFromVertex(vertex)
				term.Offset = offset
				offset += term.Length()
				resultList.PushBack(term)
				return true
			})
		}
	}else{
		for i := 0; i < length; i++ {
			resultList.Iterator(func(e *glist.Element) bool {
				var vertex Vertex = e.Value.(Vertex)
				var term Term = s.ConvertFromVertex(vertex)
				resultList.PushBack(term)
				return true
			})
		}
	}

	return *resultList
}

// 将节点转为term
func (s *Segment) ConvertFromVertex(vertex Vertex) Term {
	return *NewTerm(vertex.RealWord, vertex.GuessNature())
}
