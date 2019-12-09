package word

import (
	. "eagle/service/common"
	"eagle/utils"
	"github.com/gogf/gf/text/gstr"
)

type Word struct {
	Value string // 单词的真实值
	Label string // 单词的标签
}

func NewWord(value string, label string) *Word {
	return &Word{
		Value:value,
		Label:label,
	}
}

/**
 * 通过参数构造一个单词
 * @param param 比如 人民网/nz
 * @return 一个单词
 */
func Create(param string) *Word {
	if param == "" {
		return nil
	}
	cutIndex := gstr.PosR(param, "/")
	if cutIndex <= 0 || cutIndex == String(param).Length() - 1 {
		utils.Logger.Warning("使用 ", param, "创建单个单词失败")
		return nil
	}
	return NewWord(String(param).Substring(0, cutIndex), String(param).Substring(cutIndex+1, String(param).Length()))
}

func (w *Word) GetValue() string {
	return w.Value
}

func (w *Word) GetLabel() string {
	return w.Label
}

func (w *Word) SetLabel(label string) {
	w.Label = label
}

func (w *Word) SetValue(value string) {
	w.Value = value
}

func (w *Word) Length() int {
	return String(w.Value).Length()
}
