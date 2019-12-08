package interval

import (
	"github.com/gogf/gf/container/glist"
	"testing"
)

func TestIntervalableComparatorBySize(t *testing.T) {
	a := glist.New(true)
	for i := 10; i > 0; i-- {
		val := NewInterval(i-1, i)
		t.Log(val)
		a.PushBack(val)
	}

	t.Log(a.FrontAll())

	*a = IntervalableComparatorBySize(*a)
	for _, b := range a.FrontAll() {
		t.Log(b.(*Interval))
	}
	t.Fatal(a.FrontAll())
}
