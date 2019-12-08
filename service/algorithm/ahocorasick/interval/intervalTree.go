package interval

import (
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gset"
)

// 线段树，用于检查区间重叠
type IntervalTree struct {
	rootNode *IntervalNode // 根节点
}

// 构造线段树
func NewIntervalTree(intervals glist.List) *IntervalTree {
	iTree := &IntervalTree{}
	iTree.rootNode = NewIntervalNode(intervals)
	return iTree
}

// 从区间列表中移除重叠的区间
func (it *IntervalTree) RemoveOverlaps(intervals glist.List) glist.List {
	// 排序，按照先大小后左端点的顺序

	intervals = IntervalableComparatorBySize(intervals)

	removeIntervals := gset.NewSet(true)

	for _, iinterval := range intervals.FrontAll() {
		interval := iinterval.(Intervalable)
		// 如果区间已经被移除了，就忽略它
		if removeIntervals.Contains(interval) {
			continue
		}

		// 否则就移除它
		removeIntervals.Add(it.FindOverlaps(interval))
	}

	// 移除所有的重叠区间
	for _, removeInterval := range removeIntervals.Slice() {
		rInterval := removeInterval.(Intervalable)
		length := intervals.Len()
		for i, e := 0, intervals.Front(); i < length; i, e = i+1, e.Next() {
			if e.Value.(Intervalable) == rInterval {
				intervals.Remove(e)
				break
			}
		}
	}

	// 排序，按照左端顺序
	intervals = IntervalableComparatorByPosition(intervals)

	return intervals
}

// 寻找重叠区间
// interval 与这个区间重叠
func (it *IntervalTree) FindOverlaps(interval Intervalable) glist.List {
	return it.rootNode.FindOverlaps(interval)
}
