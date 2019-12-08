package tokenizer

import (
	. "eagle/service/common"
	"eagle/service/segment"
	"eagle/service/segment/viterbi"
	"github.com/gogf/gf/container/glist"
)

type StandardTokenizer struct {
	SEGMENT segment.ISegment
}

func NewStandardTokenizer() *StandardTokenizer {
	return &StandardTokenizer{
		SEGMENT: viterbi.NewViterbiSegment()}
}

// 分词
func (st *StandardTokenizer) Segment(text String) glist.List {
	return st.SEGMENT.Seg(text.ToCharArray())
}

// 分词
func (st *StandardTokenizer) Segment2(text []Char) glist.List {
	return st.SEGMENT.Seg(text)
}

