package interval

import (
	"github.com/gogf/gf/container/glist"
	"testing"
)

func TestNewIntervalTree(t *testing.T) {
	itree := NewIntervalTree(*glist.New(true))
	t.Fatalf("%v", itree)
}
