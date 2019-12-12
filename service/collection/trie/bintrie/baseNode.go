package bintrie

import (
	. "eagle/service/common"
	"fmt"
)

type Status int

const (
	UNDEFINED_0 Status = iota // 未指定，用于删除词条
	NOT_WORD_1 // 不是词语的结尾
	WORD_MIDDLE_2 // 是个词语的结尾，并且还可以继续
	WORD_END_3 // 是个词语的结尾，并且没有继续
)

type BaseNode struct {
	ARRAY_STATUS []Status // 状态数组，方便读取的时候用
	child []*BaseNode // 子节点
	status Status // 节点状态
	c Char // 节点代表的字符
	value interface{} // 节点代表的值
}

func NewBaseNode() *BaseNode {
	return &BaseNode{}
}

func (b *BaseNode) Transition(path string, begin int) *BaseNode {
	cur := b
	p := String(path)
	for i := begin; i < p.Length(); i++ {
		v := cur.GetChild(p.CharAt(i))
		cur = v
		if cur == nil || cur.status == UNDEFINED_0 {
			return nil
		}
	}
	return cur
}

func (b *BaseNode) Transition2(path []Char, begin int) *BaseNode {
	cur := b
	for i := begin; i < len(path); i++ {
		v := cur.GetChild(path[i])
		cur = v
		if cur == nil || cur.status == UNDEFINED_0 {
			return nil
		}
	}
	return cur
}

/**
 * 转移状态
 * @param path
 * @return
 */
func (b *BaseNode) Transition3(path Char) *BaseNode {
	cur := b
	cur = cur.GetChild(path)
	if cur == nil || cur.status == UNDEFINED_0 {
		return nil
	}
	return cur
}

/**
 * 添加子节点
 *
 * @return true-新增了节点 false-修改了现有节点
 */
func (b *BaseNode) AddChild(node BaseNode) bool {
	b.child = append(b.child, &node)
	return true
}

/**
 * 是否含有子节点
 *
 * @param c 子节点的char
 * @return 是否含有
 */
func (b *BaseNode) HasChild(c Char) bool {
	return b.GetChild(c) != nil
}

func (b *BaseNode) GetChar() Char {
	return b.c
}

/**
 * 获取子节点
 *
 * @param c 子节点的char
 * @return 子节点
 */
func (b *BaseNode) GetChild(c Char) *BaseNode {
	return NewBaseNode()
}

/**
 * 获取节点对应的值
 *
 * @return 值
 */
func (b *BaseNode) GetValue() interface{} {
	return b.value
}

/**
 * 设置节点对应的值
 *
 * @param value 值
 */
func (b *BaseNode) SetValue(value interface{}) {
	b.value = value
}

func (b *BaseNode) CompareTo2(other BaseNode) int {
	return b.CompareTo(other.GetChar())
}

func (b *BaseNode) CompareTo(other Char) int {
	if b.c > other {
		return 1
	}
	if b.c < other {
		return -1
	}
	return 0
}

/**
 * 获取节点的成词状态
 * @return
 */
func (b *BaseNode) GetStatus() Status {
	return b.status
}

func (b *BaseNode) Walk(sb String) {

}

func (b *BaseNode) ToString() string {
	if b.child == nil {
		return fmt.Sprintf("BaseNode{status=%v, c=%v, value=%v}", b.status, b.c, b.value)
	}else{
		return fmt.Sprintf("BaseNode{child=%v, status=%v, c=%v, value=%v}", len(b.child), b.status, b.c, b.value)
	}
}




