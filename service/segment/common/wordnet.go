package common

import (
	. "eagle/service/common"
	. "eagle/service/corpus/tag"
	"eagle/service/dictionary"
	"eagle/service/dictionary/coreDictionary"
	. "eagle/service/dictionary/other/charType"
	"eagle/service/segment/NShort/path"
	"eagle/service/utility/predefine"
	"github.com/gogf/gf/container/glist"
)

type WordNet struct {
	vertexes []*glist.List // 节点，每一行都是前缀词，跟图的表示方式不同
	Size int // 共有多少个节点
	Sentence String // 原始句子
	CharArray []Char // 原始句子对应的数组
}

func NewWordNet(sentence String) *WordNet {
	return NewWordNetFromCharArray(sentence.ToCharArray())
}

func NewWordNetFromCharArray(charArray []Char) *WordNet {
	wn := &WordNet{}
	wn.CharArray = charArray
	wn.vertexes = make([]*glist.List, 0, len(wn.CharArray) + 2)
	for i := 0; i< len(wn.vertexes); i++ {
		wn.vertexes[i] = glist.New(true)
	}
	wn.vertexes[0].PushBack(*NewB())
	wn.vertexes[len(wn.vertexes) - 1].PushBack(*NewE())
	wn.Size = 2
	return wn
}

// 添加顶点
func (wn *WordNet) Add(line int, vertex Vertex) {
	for _, oldVertex := range wn.vertexes[line].PopBackAll() {
		oVertex := oldVertex.(Vertex)
		if len(oVertex.RealWord) == len(vertex.RealWord) {
			return
		}
	}
	wn.vertexes[line].PushBack(vertex)
	wn.Size++
}

// 添加顶点，由原子分词顶点添加
func (wn *WordNet) AddSegment(line int, atomSegment glist.List) {
	offset := 0
	for _, atomNode := range atomSegment.FrontAll() {
		aNode := atomNode.(path.AtomNode)
		sWord := aNode.SWord
		nature := GNature.N
		id := -1
		switch aNode.NPOS {
		case GCharType.CT_CHINESE.ToInt():
		case GCharType.CT_NUM.ToInt(): fallthrough
		case GCharType.CT_INDEX.ToInt(): fallthrough
		case GCharType.CT_CNUM.ToInt():
			nature = GNature.N
			sWord = predefine.TAG_NUMBER
			id = coreDictionary.GCoreDictionary.M_WORD_ID
		case GCharType.CT_DELIMITER.ToInt(): fallthrough
		case GCharType.CT_OTHER.ToInt():
			nature = GNature.W
		case GCharType.CT_SINGLE.ToInt():
			nature = GNature.Nx
			sWord = predefine.TAG_CLUSTER
			id = coreDictionary.GCoreDictionary.X_WORD_ID
		default:
		}
		// 这些通用符的量级都在10万左右
		wn.Add(line + offset, *NewVertex2(sWord, aNode.SWord, *dictionary.NewAttribute3(*nature, 10000), id))
		offset += len(aNode.SWord)
	}
}

// 获取某一行的所有节点
func (wn *WordNet) Get(line int) glist.List {
	return *wn.vertexes[line]
}

func (wn *WordNet) GetVertexes() []*glist.List {
	return wn.vertexes
}

