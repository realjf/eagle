package trie

import (
	"fmt"
	"eagle/service/algorithm/ahocorasick/interval"
)

// 一个模式串匹配结果
type Emit struct {
	keyword string // 匹配到的模式串
	interval.Interval
}

// 构造一个模式串匹配结果
// start 起点
// end 重点
// keyword 模式串
func NewEmit(start int, end int, keyword string) *Emit {
	return &Emit{
		Interval: *interval.NewInterval(start, end),
		keyword:  keyword,
	}
}

// 获取对应的模式串
func (e *Emit) GetKeyword() string {
	return e.keyword
}

func (e *Emit) ToString() string {
	return fmt.Sprintf("%s=%s", e.Interval.ToString(), e.keyword)
}
