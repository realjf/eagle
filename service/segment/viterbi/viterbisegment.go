package viterbi

import (
	"eagle/service/collection/trie"
	. "eagle/service/common"
	"eagle/service/config"
	"eagle/service/segment"
	"eagle/service/segment/common"
	"eagle/utils"
	"github.com/gogf/gf/container/glist"
)

type ViterbiSegment struct {
	dat *trie.DoubleArrayTrie
	segment.WordBasedSegment
}

func NewViterbiSegment() *ViterbiSegment {
	return &ViterbiSegment{
		dat:
	}
}

// 切分句子
func (s *ViterbiSegment) SegSentence(sentence []Char) glist.List {
	var wordNetAll *common.WordNet = common.NewWordNetFromCharArray(sentence)
	////////////////生成词网////////////////////
	s.GenerateWordNet(wordNetAll)
	if config.GConfig.DEBUG {
		utils.Logger.Printf("粗分词网：\n%s\n", wordNetAll)
	}
	///////////////生成词图////////////////////
	vertexList := s.viterbi(wordNetAll)
	if s.Config.UseCustomDictionary {
		if s.Config.IndexMode > 0 {

		}else{

		}
	}

	if config.GConfig.DEBUG {
		utils.Logger.Println("粗分结果", s.Convert(vertexList, false))
	}

	return *glist.New(true)
}

func (s *ViterbiSegment) viterbi(wordNet *common.WordNet) glist.List {
	nodes := wordNet.GetVertexes()
	vertexList := glist.New(true)
	nodes[1].Iterator(func(e *glist.Element) bool {
		nv := e.Value.(common.Vertex)
		nv.UpdateFrom(nodes[0].FrontValue().(common.Vertex))
		return true
	})
	for i := 0; i < len(nodes) - 1; i++ {
		nodeArray := nodes[i]
		if nodeArray == nil {
			continue
		}
		for _, node := range nodeArray.FrontAll() {
			nv := node.(common.Vertex)
			if nv.From == nil {
				continue
			}
			nodes[i+String(nv.RealWord).Length()].Iterator(func(e *glist.Element) bool {
				to := e.Value.(common.Vertex)
				to.UpdateFrom(nv)
				return true
			})
		}
	}
	var from common.Vertex = nodes[len(nodes) - 1].FrontValue().(common.Vertex)
	for &from != nil {
		vertexList.PushFront(from)
		from = *from.From
	}
	return *vertexList
}
