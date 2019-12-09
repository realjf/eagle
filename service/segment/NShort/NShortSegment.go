package NShort

import (
	. "eagle/service/common"
	"eagle/service/segment"
	. "eagle/service/segment/common"
	"github.com/gogf/gf/container/glist"
)

type NShortSegment struct {
	segment.WordBasedSegment
}

func NewNShortSegment() *NShortSegment {
	return &NShortSegment{}
}

func (ns *NShortSegment) SegSentence(sentence []Char) glist.List {

}

/**
 * 二元语言模型分词
 * @param sSentence 待分词的句子
 * @param nKind     需要几个结果
 * @param wordNetOptimum
 * @param wordNetAll
 * @return 一系列粗分结果
 */
func (ns *NShortSegment) BiSegment(sSentence []Char, nKind int, wordNetOptimum WordNet, wordNetAll WordNet) glist.List {

}
