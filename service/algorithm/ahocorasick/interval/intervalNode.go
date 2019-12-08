package interval

import "github.com/gogf/gf/container/glist"

type DIRECTION int

const (
	LEFT DIRECTION = iota
	RIGHT
)

// 线段树上面的节点，实际上是一些区间的集合，并且按中点维护了两个节点
type IntervalNode struct {
	left      *IntervalNode // 区间集合的最左端
	right     *IntervalNode // 区间集合的最右端
	point     int           // 中点
	intervals glist.List    // 区间集合

}

func NewIntervalNode(intervals glist.List) *IntervalNode {
	intervalNode := &IntervalNode{}
	intervalNode.point = intervalNode.DetermineMedian(intervals)
	toLeft := glist.New(true)  // 以中点为界靠左的区间
	toRight := glist.New(true) // 靠右的区间

	for _, iinterval := range intervals.FrontAll() {
		interval := iinterval.(Intervalable)
		if interval.GetEnd() < intervalNode.point {
			toLeft.PushBack(interval)
		} else if interval.GetStart() > intervalNode.point {
			toRight.PushBack(interval)
		} else {
			intervalNode.intervals.PushBack(interval)
		}
	}

	if toLeft.Size() > 0 {
		intervalNode.left = NewIntervalNode(*toLeft)
	}
	if toRight.Size() > 0 {
		intervalNode.right = NewIntervalNode(*toRight)
	}

	return intervalNode
}

// 计算中点
// intervals 区间集合
func (in *IntervalNode) DetermineMedian(intervals glist.List) int {
	start := -1
	end := -1
	for _, iinterval := range intervals.FrontAll() {
		interval := iinterval.(Intervalable)
		currentStart := interval.GetStart()
		currentEnd := interval.GetEnd()
		if start == -1 || currentStart < start {
			start = currentStart
		}
		if end == -1 || currentEnd > end {
			end = currentEnd
		}
	}
	return (start + end) / 2
}

// 寻找与interval有重叠的区间
func (in *IntervalNode) FindOverlaps(interval Intervalable) glist.List {
	overlaps := glist.New(true)

	if in.point < interval.GetStart() {
		// 右边找找
		in.AddToOverlaps(interval, *overlaps, in.FindOverlappingRanges(*in.right, interval))
		in.AddToOverlaps(interval, *overlaps, in.CheckForOverlapsToTheRight(interval))
	} else if in.point > interval.GetEnd() {
		// 左边找找
		in.AddToOverlaps(interval, *overlaps, in.FindOverlappingRanges(*in.left, interval))
		in.AddToOverlaps(interval, *overlaps, in.CheckForOverlapsToTheLeft(interval))
	} else {
		// 否则在当前区间
		in.AddToOverlaps(interval, *overlaps, in.intervals)
		in.AddToOverlaps(interval, *overlaps, in.FindOverlappingRanges(*in.left, interval))
		in.AddToOverlaps(interval, *overlaps, in.FindOverlappingRanges(*in.right, interval))
	}
	return *overlaps
}

// 添加到重叠区间列表中
// interval 跟此区间重叠
// overlaps 重叠区间列表
// newOverlaps 希望将这些区间加入
func (in *IntervalNode) AddToOverlaps(interval Intervalable, overlaps glist.List, newOverlaps glist.List) {
	for _, currentInterval := range newOverlaps.FrontAll() {
		curInterval := currentInterval.(Intervalable)
		if curInterval != interval {
			overlaps.PushBack(curInterval)
		}
	}
}

// 往左边寻找重叠
func (in *IntervalNode) CheckForOverlapsToTheLeft(interval Intervalable) glist.List {
	return in.CheckForOverlaps(interval, LEFT)
}

// 往右边寻找重叠
func (in *IntervalNode) CheckForOverlapsToTheRight(interval Intervalable) glist.List {
	return in.CheckForOverlaps(interval, RIGHT)
}

// 寻找重叠
// interval 一个区间，与该区间重叠
// direction 方向，表明重叠区间在interval的左边还是右边
func (in *IntervalNode) CheckForOverlaps(interval Intervalable, direction DIRECTION) glist.List {
	overlaps := glist.New(true)
	for _, currentInterval := range in.intervals.FrontAll() {
		curInterval := currentInterval.(Intervalable)
		switch direction {
		case LEFT:
			if curInterval.GetStart() <= interval.GetEnd() {
				overlaps.PushBack(curInterval)
			}
		case RIGHT:
			if curInterval.GetEnd() >= interval.GetStart() {
				overlaps.PushBack(curInterval)
			}
		}
	}
	return *overlaps
}

// 是对IntervalNode.findOverlaps(Intervalable)的一个包装，防止NPE
func (in *IntervalNode) FindOverlappingRanges(node IntervalNode, interval Intervalable) glist.List {
	if &node == nil {
		return node.FindOverlaps(interval)
	}
	return glist.List{}
}
