package ahocorasick

import "fmt"

//命中一个模式串的处理方法
type IHit interface {
	// begin 模式串在母文本中的起始位置
	// end   模式串在母文本中的终止位置
	// value 模式串对应的值
	Hit(begin int, end int, value interface{})
}

type IHitFull interface {
	// begin 模式串在母文本中的起始位置
	// end   模式串在母文本中的终止位置
	// value 模式串对应的值
	// index 模式串对应的值的下标
	Hit(begin int, end int, value interface{}, index int)
}

// 一个命中的结果
type Hit struct {
	begin int         // 模式串在母文本中的起始位置
	end   int         // 模式串在母文本中的终止位置
	value interface{} // 模式串对应的值
}

func NewHit(begin int, end int, value interface{}) *Hit {
	return &Hit{
		begin:begin,
		end:end,
		value: value,
	}
}

func (h *Hit) Hit(begin int, end int, value interface{}) {
	h.begin = begin
	h.end = end
	h.value = value
}

func (h *Hit) ToString() string {
	return fmt.Sprintf("[%d:%d]=%s", h.begin, h.end, h.value)
}
