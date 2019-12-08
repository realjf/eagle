package interval

import (
	"github.com/gogf/gf/container/glist"
)

// 按照长度比较区间，如果长度相同，则比较起点
func IntervalableComparatorBySize(a glist.List) glist.List {
	b := glist.New(true)

	length := a.Len()
	for i := 0; i < length; i++ {
		gard := a.PopFront()
		len := a.Len()
		if len > 0 {
			for i := 0; i < len; i++ {
				e := a.PopFront()
				if e == nil {
					break
				}
				intervalable1 := gard.(Intervalable)
				intervalable2 := e.(Intervalable)
				comparision := intervalable2.Size() - intervalable1.Size()
				if comparision == 0 {
					// 长度相同，比较起点
					comparision = intervalable1.GetStart() - intervalable2.GetStart()
				}
				if comparision > 0 {
					a.PushBack(gard)
					gard = e
				} else {
					a.PushBack(intervalable2)
				}
			}
		}

		b.PushBack(gard)
	}

	return *b
}
