package common

import (
	"eagle/service/corpus/tag"
)

type Term struct {
	Word string // 词语
	Nature tag.Nature // 词性
	Offset int // 在文本中的起始位置（需开启分词器的offset选项）

}

func NewTerm(word string, nature tag.Nature) *Term {
	return &Term{
		Word: word,
		Nature: nature,
	}
}

func (t *Term) ToString() string {
	// 判断是否显示词性

	return t.Word
}

// 长度
func (t *Term) Length() int {
	return len(t.Word)
}

// 获取本词语在HanLP词库中的频次
func (t *Term) GetFrequency() int {
	return 0
}

// 判断Term是否相等
func (t *Term) Equals(obj interface{}) bool {
	if o, ok := obj.(Term); ok {
		if t.Nature.ToString() == o.Nature.ToString() && t.Word == o.Word {
			return true
		}
	}
	return false
}


