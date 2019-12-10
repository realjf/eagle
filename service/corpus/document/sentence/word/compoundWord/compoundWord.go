package compoundWord

import (
	"eagle/service/common"
	"eagle/service/corpus/document/sentence/word"
	"eagle/utils"
	"github.com/gogf/gf/container/glist"
	"strings"
)

// 复合词，由两个或以上的word构成
type CompoundWord struct {
	InnerList glist.List // 由这些词复合而来
	Label string // 标签，通常是词性

}

func (cw *CompoundWord) GetValue() string {
	sb := common.NewString()
	for _, w := range cw.InnerList.FrontAll() {
		sb.Append(w.(word.Word).Value)
	}
	return sb.ToString()
}

func (cw *CompoundWord) GetLabel() string {
	return cw.Label
}

func (cw *CompoundWord) SetLabel(label string) {
	cw.Label = label
}

func (cw *CompoundWord) SetValue(value string) {
	cw.InnerList.Clear()
	cw.InnerList.PushBack(word.NewWord(value, cw.Label))
}

func (cw *CompoundWord) Length() int {
	return common.String(cw.GetValue()).Length()
}

func (cw *CompoundWord) ToString() string {
	sb := common.NewString()
	sb.Append('[')
	i := 1
	for _, w := range cw.InnerList.FrontAll() {
		sb.Append(w.(word.Word).GetValue())
		label := w.(word.Word).GetLabel()
		if label != "" {
			sb.Append('/').Append(label)
		}
		if i != cw.InnerList.Size() {
			sb.Append(' ')
		}
		i++
	}
	sb.Append("]/")
	sb.Append(cw.Label)
	return sb.ToString()
}

func (cw *CompoundWord) ToWord() word.Word {
	return *word.NewWord(cw.GetValue(), cw.GetLabel())
}

func NewCompoundWord(innerList glist.List, label string) *CompoundWord {
	cw := &CompoundWord{
		InnerList:innerList,
		Label:label,
	}
	return cw
}

func Create(param string) *CompoundWord {
	if param == "" {
		return nil
	}
	params := common.String(param)
	cutIndex := params.LastIndexOf("]")
	if cutIndex <= 2 || cutIndex == params.Length() -1 {
		return nil
	}
	wordParam := params.Substring(1, cutIndex)
	wordList := glist.New(true)
	for _, single := range strings.Split(wordParam, " ") {
		if len(single) == 0 {
			continue
		}
		w := word.Create(single)
		if w == nil {
			utils.Logger.Warning("使用参数", single, "构造单词时发生错误")
			return nil
		}
		wordList.PushBack(w)
	}
	labelParam := params.Substring(cutIndex+1, params.Length())
	if common.String(labelParam).StartsWith("/") {
		labelParam = common.String(labelParam).Substring(1, common.String(labelParam).Length())
	}
	return NewCompoundWord(*wordList, labelParam)
}
