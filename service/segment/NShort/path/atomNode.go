package path

import (
	"eagle/service/dictionary"
	"eagle/service/dictionary/other/charType"
	"fmt"
	. "eagle/service/common"
	. "eagle/service/corpus/tag"
	"eagle/service/utility/predefine"
	"eagle/service/segment/common"
	. "eagle/service/utility/textUtility"
)

// 原子分词节点
type AtomNode struct {
	SWord string
	NPOS int

}

func NewAtomNode(sWord string, nPOS int) *AtomNode {
	return &AtomNode{
		SWord:sWord,
		NPOS:nPOS,
	}
}

func NewAtomNode2(c Char, nPOS int) *AtomNode {
	return &AtomNode{
		SWord: c.ToString(),
		NPOS:nPOS,
	}
}

func (an *AtomNode) GetNature() Nature {
	nature := NewNature("").Nz
	switch an.NPOS {
	case CT_CHINESE:
	case CT_NUM: fallthrough
	case CT_INDEX: fallthrough
	case CT_CNUM:
		nature = NewNature("").M
		an.SWord = predefine.TAG_NUMBER
	case CT_DELIMITER:
		nature = NewNature("").W
	case CT_LETTER:
		nature = NewNature("'").Nx
		an.SWord = predefine.TAG_CLUSTER
	case CT_SINGLE:
		if predefine.PATTERN_FLOAT_NUMBER.MatchString(an.SWord) {//匹配浮点数
			nature = NewNature("").M
			an.SWord = predefine.TAG_NUMBER
		}else{
			nature = NewNature("").Nx
			an.SWord = predefine.TAG_CLUSTER
		}
	default:

	}

	return *nature
}

func (an *AtomNode) ToString() string {
	return fmt.Sprintf("AtomNode{word='%s', nature=%d", an.SWord, an.NPOS)
}


func (an *AtomNode) Convert(word string, t int) common.Vertex {
	name := word
	nature := GNature.N
	dValue := 1

	switch t {
	case charType.GCharType.CT_CHINESE.ToInt():
	case charType.GCharType.CT_NUM.ToInt():fallthrough
	case charType.GCharType.CT_INDEX.ToInt():fallthrough
	case charType.GCharType.CT_CNUM.ToInt():
		nature = GNature.M
		word = predefine.TAG_NUMBER
	case charType.GCharType.CT_DELIMITER.ToInt():
		nature = GNature.W
	case charType.GCharType.CT_LETTER.ToInt():
		nature = GNature.Nx
		word = predefine.TAG_CLUSTER
	case charType.GCharType.CT_SINGLE.ToInt():
		nature = GNature.Nx
		word = predefine.TAG_CLUSTER
	}

	return *common.NewVertex(word, name, *dictionary.NewAttribute3(*nature, dValue))
}

