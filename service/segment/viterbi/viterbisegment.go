package viterbi

import (
	. "eagle/service/common"
	"eagle/service/segment"
	"eagle/service/segment/common"
	"github.com/gogf/gf/container/glist"
)

type ViterbiSegment struct {
	segment.WordBasedSegment
}

func NewViterbiSegment() *ViterbiSegment {
	return &ViterbiSegment{

	}
}

// 切分句子
func (s *ViterbiSegment) SegSentence(sentence []Char) glist.List {
	wordNetAll := common.NewWordNetFromCharArray(sentence)



	return *glist.New(true)
}
