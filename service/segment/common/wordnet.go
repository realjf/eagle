package common

import (
	. "eagle/service/common"
	"github.com/gogf/gf/container/glist"
)

type WordNet struct {
	vertexes []*glist.List // 节点，每一行都是前缀词，跟图的表示方式不同
	Size int // 共有多少个节点
	Sentence String // 原始句子
	charArray []Char // 原始句子对应的数组
}

func NewWordNet(sentence String) *WordNet {
	return NewWordNetFromCharArray(sentence.ToCharArray())
}

func NewWordNetFromCharArray(charArray []Char) *WordNet {
	wn := &WordNet{}
	wn.charArray = charArray
	wn.vertexes = make([]*glist.List, 0, len(wn.charArray) + 2)
	for i := 0; i< len(wn.vertexes); i++ {
		wn.vertexes[i] = glist.New(true)
	}
	wn.vertexes[0].PushBack(Vertex)
}
