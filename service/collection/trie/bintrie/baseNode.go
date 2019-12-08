package bintrie

import . "eagle/service/common"

type Status int

const (
	UNDEFINED_0 Status = iota // 未指定，用于删除词条
	NOT_WORD_1 // 不是词语的结尾
	WORD_MIDDLE_2 // 是个词语的结尾，并且还可以继续
	WORD_END_3 // 是个词语的结尾，并且没有继续
)

type BaseNode struct {
	ARRAY_STATUS []Status // 状态数组，方便读取的时候用
	child []BaseNode // 子节点
	status Status // 节点状态
	c Char // 节点代表的字符
	value interface{} // 节点代表的值
}




