package word

// 词语接口
type IWord interface {
	GetValue() string // 获取单词
	GetLabel() string // 获取标签
	SetLabel(label string) // 设置标签
	SetValue(value string) // 设置单词
	Length() int // 单词长度
}
