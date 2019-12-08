package common

import "fmt"

// 记录了起点的边
type EdgeFrom struct {
	Edge
	From int
}

func NewEdgeFrom(from int, weight float64, name string) *EdgeFrom {
	return &EdgeFrom{
		From: from,
		Edge: Edge{
			Weight:weight,
			Name:name,
		},
	}
}

func (ef *EdgeFrom) ToString() string {
	return fmt.Sprintf("EdgeFrom{ from=%d, weight= %f, name='%s'}", ef.From, ef.Weight, ef.Name)
}