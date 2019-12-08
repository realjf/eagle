package customDictionary

import (
	"eagle/service/collection/trie"
	"eagle/service/collection/trie/bintrie"
)

var (
	GCustomDictionary *CustomDictionary
)

func init()  {
	if GCustomDictionary == nil {
		GCustomDictionary = NewCustomDictionary()
	}
}

type CustomDictionary struct {
	Trie *bintrie.BinTrie
	Dat *trie.DoubleArrayTrie
}

func NewCustomDictionary() *CustomDictionary {
	return &CustomDictionary{
		Dat: trie.NewDoubleArrayTrie(),
	}
}
