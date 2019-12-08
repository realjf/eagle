package algorithm

import (
	"eagle/service/common"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/glist"
)

// 用固定容量的优先队列模拟的最大堆，用于解决求topN大的问题
type MaxHeap struct {
	queue common.PriorityQueue // 优先队列
	maxSize int // 堆的最大容量

}

// 构造最大堆
// maxSize 保留多少个元素
func NewMaxHeap(maxSize int) *MaxHeap {
	if maxSize <= 0 {
		return nil
	}
	mh := &MaxHeap{
		maxSize:maxSize,
		queue: common.NewPriorityQueue(maxSize),
	}

	return mh
}

// 添加一个元素
func (mh *MaxHeap) Add(e interface{}, priority int) bool {
	if mh.queue.Len() < mh.maxSize {
		item := common.Item{Priority: priority, Value: e}
		mh.queue.Push(&item)
		return true
	}else{
		// 队列已满
		peek, _ := mh.queue.Peek(0)
		if peek.Priority > priority {
			// 将新元素与当前堆顶元素比较，保留较小的元素
			item := common.Item{Priority:priority,Value:e}
			mh.queue.Poll()
			mh.queue.Push(item)
			return true
		}
	}
	return false
}

// 添加许多元素
func (mh *MaxHeap) AddAll(collection garray.Array) *MaxHeap {
	for _, v := range collection.Slice() {
		mh.Add(v, 0)
	}
	return mh
}

// 转为有序列表，自毁性操作
func (mh *MaxHeap) ToList() glist.List {
	list := glist.New(true)
	for mh.queue.Len() != 0 {
		item, _ := mh.queue.Poll()
		if item != nil {
			list.PushBack(*item)
		}
	}
	return *list
}

func (mh *MaxHeap) Size() int {
	return mh.queue.Len()
}


