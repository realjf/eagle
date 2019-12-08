package trie

import (
	"eagle/service/common"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"io"
)

type ITrie interface {
	Build(keyValueMap gmap.TreeMap) int
	Save(writer io.Writer) bool
	Load(array garray.Array, value interface{}) bool
	Get(key []common.Char) interface{}
	Get2(key string) interface{}
	GetValueArray(a []interface{}) []interface{}
	ContainsKey(key string) bool
	Size() int
}
