package path

type QueueElement struct {
	From int // 边的起点
	Index int // 边的终点在顶点数组中的下标
	Weight float64
	Next *QueueElement
}

func NewQueueElement(from int, index int, weight float64) *QueueElement {
	return &QueueElement{
		From:from,
		Index:index,
		Weight:weight,
	}
}

func (qe *QueueElement) CompareTo(other QueueElement) int {
	if qe.Weight < other.Weight {
		return -1
	}else if qe.Weight > qe.Weight {
		return 1
	}else{
		return 0
	}
}
