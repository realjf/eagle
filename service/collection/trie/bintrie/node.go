package bintrie

import (
	"eagle/service/collection/trie/bintrie/util"
	. "eagle/service/common"
)

type Node struct {
	BaseNode
}

/**
 * @param c      节点的字符
 * @param status 节点状态
 * @param value  值
 */
func NewNode(c Char, status Status, value interface{}) *Node {
	return &Node{
		BaseNode{
			c:c,
			status:status,
			value:value,
		},
	}
}

func (n *Node) AddChild(node BaseNode) bool {
	add := false
	if n.child == nil {
		n.child = make([]*BaseNode, 0)
	}
	index := util.BinarySearch(n.child, node)
	if index >= 0 {
		target := n.child[index]
		switch node.status {
		case UNDEFINED_0:
			if target.status != NOT_WORD_1 {
				target.status = NOT_WORD_1
				target.value = nil
				add = true
			}
		case NOT_WORD_1:
			if target.status == WORD_END_3 {
				target.status = WORD_END_3
			}
		case WORD_END_3:
			if target.status != WORD_END_3 {
				target.status = WORD_MIDDLE_2
			}
			if target.GetValue() == nil {
				add = true
			}
			target.SetValue(node.GetValue())
		}
	}else {
		newChild := make([]*BaseNode, len(n.child))
		insert := -(index + 1)
		copy(newChild, n.child[0:insert])
		copy(newChild[insert+1:], n.child[insert:len(n.child) - insert])
		newChild[insert] = &node
		n.child = newChild
		add = true
	}
	return add
}

func (n *Node) GetChild(c Char) *BaseNode {
	if n.child == nil {
		return nil
	}
	var index int = util.BinarySearch2(n.child, c)
	if index < 0 {
		return nil
	}
	return n.child[index]
}

