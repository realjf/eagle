package common

import "container/heap"

type Item struct {
	Value interface{}
	Index int
	Priority int
}

// 优先队列
type PriorityQueue []*Item

func NewPriorityQueue(cap int) PriorityQueue {
	return make(PriorityQueue, 0, cap)
}

// 实现heap.Interface接口
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// 添加新元素
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(PriorityQueue, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	*pq = (*pq)[0:n+1]
	item := x.(*Item)
	item.Index = n
	(*pq)[n] = item
}

// 移除元素并返回元素
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	c := len(*pq)
	if n < (c/2) && c > 25 {
		npq := make(PriorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	item := (*pq)[n-1]
	item.Index = -1
	*pq = (*pq)[0:n-1]
	return item
}

func (pq *PriorityQueue) PeekAndShift(max int) (*Item, int) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)
	return item, 0
}

// 移除队首元素
func (pq *PriorityQueue) Poll() (*Item, int) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	heap.Remove(pq, 0)
	return item, 0
}

func (pq *PriorityQueue) Peek(max int) (*Item, int) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}

	return item, 0
}
