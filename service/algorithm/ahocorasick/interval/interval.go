package interval

import "fmt"

// 区间
type Interval struct {
	start int // 起点
	end   int // 终点
}

// 构造一个区间
func NewInterval(start int, end int) *Interval {
	return &Interval{
		start: start,
		end:   end,
	}
}

func (in *Interval) GetStart() int {
	return in.start
}

func (in *Interval) GetEnd() int {
	return in.end
}

func (in *Interval) Size() int {
	return in.end - in.start + 1
}

// 是否与另一个区间交叉（有一部分重叠）
func (in *Interval) OverlapsWith(other Interval) bool {
	return in.start <= other.GetEnd() && in.end >= other.GetStart()
}

// 区间是否覆盖了这个点
func (in *Interval) OverlapsWith2(point int) bool {
	return in.start <= point && point <= in.end
}

func (in *Interval) Equals(obj interface{}) bool {
	if o, ok := obj.(Intervalable); !ok {
		return false
	} else {
		return in.start == o.GetStart() && in.end == o.GetEnd()
	}
}

func (in *Interval) HashCode() int {
	return in.start%100 + in.end%100
}

func (in *Interval) CompareTo(obj interface{}) int {
	if o, ok := obj.(Intervalable); !ok {
		return -1
	} else {
		comparison := in.start - o.GetStart()
		if comparison != 0 {
			return comparison
		} else {
			return in.end - o.GetEnd()
		}
	}
}

func (in *Interval) ToString() string {
	return fmt.Sprintf("%d:%d", in.start, in.end)
}
