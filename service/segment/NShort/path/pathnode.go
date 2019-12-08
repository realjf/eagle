package path

import "fmt"

// 路径上的节点
type PathNode struct {
	From int // 节点前驱
	Index int // 节点在顶点数组中的下标
}

/**
 * 构造一个节点
 * @param from 节点前驱
 * @param index 节点在顶点数组中的下标
 */
func NewPathNode(from int, index int) *PathNode {
	return &PathNode{
		From:  from,
		Index: index,
	}
}

func (pn *PathNode) ToString() string {
	return fmt.Sprintf("PathNode{from=%d, index=%d}", pn.From, pn.Index)
}
