package common

// 基础边，不允许构造
type Edge struct {
	Weight float64 // 花费
	Name string // 节点名字，调试用
}

func NewEdge(weight float64, name string) *Edge {
	return &Edge{
		Weight: weight,
		Name: name,
	}
}

