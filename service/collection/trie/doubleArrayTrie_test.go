package trie

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
	"testing"
)

func TestNewDoubleArrayTrie(t *testing.T) {

}

func TestHandleEmptyString(t *testing.T) {
	emptyString := ""
	dat := NewDoubleArrayTrie()
	dict := gmap.NewTreeMap(gutil.ComparatorString, true)
	dict.Set("bug", "问题")
	dat.Build4(*dict)
	searcher := dat.GetSearcher2(emptyString, 0)
	for searcher.Next() {
		t.Log(searcher)
	}

	t.FailNow()
}
