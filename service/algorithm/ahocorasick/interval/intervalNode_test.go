package interval

import (
	"github.com/gogf/gf/container/glist"
	"testing"
)

func TestNewIntervalNode(t *testing.T) {
	iNode := NewIntervalNode(*glist.New(true))
	t.Fatalf("%v", iNode)
}
