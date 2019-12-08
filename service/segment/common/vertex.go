package common

import (
	. "eagle/service/corpus/tag"
	"eagle/service/corpus/tag"
	"eagle/service/dictionary"
	"eagle/service/dictionary/coreDictionary"
	"eagle/service/utility/mathUtility"
	"eagle/service/utility/predefine"
	"eagle/utils"
)

// 顶点
type Vertex struct {
	Word string // 节点对应的词或等效词（如未##数）
	RealWord string // 节点对应的真实词，绝对不含##
	Attribute dictionary.Attribute // 词的属性，谨慎修改属性内部的数据，因为会影响到字典
	WordID int // 等效词ID,也是Attribute的下标
	Index int // 在一维顶点数组中的下标，可以视作这个顶点的id

	From *Vertex // 到该节点的最短路径的前驱节点
	Weight float64 // 最短路径对应的权重
}

func NewVertex(word string, realWord string, attribute dictionary.Attribute) *Vertex {
	var wordID int
	if &attribute == nil {
		wordID = -1
	}else{
		wordID = attribute.TotalFrequency
	}
	return NewVertex2(word, realWord, attribute, wordID)
}

func NewVertex2(word string, realWord string, attribute dictionary.Attribute, wordID int) *Vertex {
	if &attribute == nil {
		attribute = *dictionary.NewAttribute3(*NewNature("").N, 1)
	}
	v := &Vertex{
		RealWord: realWord,
		WordID: wordID,
		Attribute:attribute,
	}
	if word == "" {
		word = v.CompileRealWord(realWord, attribute)
	}
	v.Word = word
	return v
}

// 生成线程安全的起始节点
func NewB() *Vertex {
	return NewVertex2(predefine.TAG_BIGIN, " ", *dictionary.NewAttribute3(*NewNature("").Begin, predefine.MAX_FREQUENCY / 10), coreDictionary.GCoreDictionary.GetWordID(predefine.TAG_BIGIN))
}

func (v *Vertex) CompileRealWord(realWord string, attribute dictionary.Attribute) string {
	if len(attribute.Nature) == 1 {
		nature := attribute.Nature[0]
		if nature.StartsWith("nr") {
			v.WordID = coreDictionary.GCoreDictionary.NR_WORD_ID
			return predefine.TAG_PEOPLE
		}else if nature.StartsWith("ns") {
			v.WordID = coreDictionary.GCoreDictionary.NS_WORD_ID
			return predefine.TAG_PLACE
		}else if nature.ToString() == NewNature("").Nx.ToString() {
			v.WordID = coreDictionary.GCoreDictionary.NX_WORD_ID
			if v.WordID == -1 {
				v.WordID = coreDictionary.GCoreDictionary.X_WORD_ID
			}
			return predefine.TAG_PEOPLE
		}else if nature.StartsWith("nt") || nature.ToString() == NewNature("").Nit.ToString() {
			v.WordID = coreDictionary.GCoreDictionary.NT_WORD_ID
			return predefine.TAG_GROUP
		}else if nature.StartsWith("m") {
			v.WordID = coreDictionary.GCoreDictionary.M_WORD_ID
			v.Attribute = coreDictionary.GCoreDictionary.GetByWordID(coreDictionary.GCoreDictionary.M_WORD_ID)
			return predefine.TAG_NUMBER
		}else if nature.StartsWith("x") {
			v.WordID = coreDictionary.GCoreDictionary.X_WORD_ID
			v.Attribute = coreDictionary.GCoreDictionary.GetByWordID(coreDictionary.GCoreDictionary.X_WORD_ID)
			return predefine.TAG_CLUSTER
		}else if nature.ToString() == NewNature("").T.ToString() {
			v.WordID = coreDictionary.GCoreDictionary.T_WORD_ID
			v.Attribute = coreDictionary.GCoreDictionary.GetByWordID(coreDictionary.GCoreDictionary.T_WORD_ID)
			return predefine.TAG_TIME
		}
	}

	return realWord
}

func (v *Vertex) UpdateFrom(from Vertex) {
	var weight float64 = from.Weight + mathUtility.CalculateWeight(from, *v)
	if v.From == nil || v.Weight > weight {
		v.From = &from
		v.Weight = weight
	}
}

func (v *Vertex) GetRealWord() string {
	return v.RealWord
}

func (v *Vertex) GetFrom() Vertex {
	return *v.From
}

func (v *Vertex) SetFrom(from Vertex) {
	v.From = &from
}

// 获取词的属性
func (v *Vertex) GetAttribute() dictionary.Attribute {
	return v.Attribute
}

func (v *Vertex) ConfirmNature(nature tag.Nature) bool {

	result := true

	return result
}

// 将属性的词性锁定为nature，此重载会降低性能
// nature     词性
// updateWord 是否更新预编译字串
func (v *Vertex) ConfirmNature2(nature Nature, updateWord bool) bool {
	switch nature.FirstChar() {
	case "m":
		v.Word = predefine.TAG_NUMBER
	case "t":
		v.Word = predefine.TAG_TIME
	default:
		// 记录错误
		utils.Logger.Warning("没有与" + nature.ToString() + " 对应的case")
	}

	return v.ConfirmNature(nature)
}

// 获取该节点的词性，如果词性还未确定，则返回nil
func (v *Vertex) GetNature() Nature {
	if len(v.Attribute.Nature) == 1 {
		return v.Attribute.Nature[0]
	}
	return Nature{}
}

// 猜测最可能的词性，也就是这个节点的词性中出现频率最大的那一个词性
func (v *Vertex) GuessNature() Nature {
	return v.Attribute.Nature[0]
}

func (v *Vertex) HasNature(nature Nature) bool {
	return v.Attribute.GetNatureFrequencyFromNature(nature) > 0
}

func (v *Vertex) SetWord(word string) Vertex {
	v.Word = word
	return *v
}

func (v *Vertex) SetRealWord(realWord string) Vertex {
	v.RealWord = realWord
	return *v
}

func (v *Vertex) Length() int {
	return len(v.RealWord)
}

func (v *Vertex) ToString() string {
	return v.RealWord
}