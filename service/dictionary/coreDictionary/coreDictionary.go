package coreDictionary

import (
	"eagle/service/collection/trie"
	"eagle/service/config"
	"eagle/service/dictionary"
	"eagle/service/utility/predefine"
)

const (
	totalFrequency = 221894
)

var (
	GCoreDictionary *CoreDictionary
)

func init() {
	if GCoreDictionary == nil {
		GCoreDictionary = NewCoreDictionary()
	}
}

// 使用DoubleArrayTrie实现的核心词典
type CoreDictionary struct {
	Path string
	Trie *trie.DoubleArrayTrie
	TotalFrequency int

	NR_WORD_ID int
	NS_WORD_ID int
	NT_WORD_ID int
	T_WORD_ID int
	X_WORD_ID int
	M_WORD_ID int
	NX_WORD_ID int
}

func NewCoreDictionary() *CoreDictionary {
	cd := &CoreDictionary{
		Trie: trie.NewDoubleArrayTrie(),
		Path: config.GConfig.CoreDictionaryPath,
		TotalFrequency: totalFrequency,
	}

	cd.NR_WORD_ID = cd.GetWordID(predefine.TAG_PEOPLE)
	cd.NS_WORD_ID = cd.GetWordID(predefine.TAG_PLACE)
	cd.NT_WORD_ID = cd.GetWordID(predefine.TAG_GROUP)
	cd.T_WORD_ID = cd.GetWordID(predefine.TAG_TIME)
	cd.X_WORD_ID = cd.GetWordID(predefine.TAG_CLUSTER)
	cd.M_WORD_ID = cd.GetWordID(predefine.TAG_NUMBER)
	cd.NX_WORD_ID = cd.GetWordID(predefine.TAG_PEOPLE)

	return cd
}

func (cd *CoreDictionary) Get(key string) dictionary.Attribute {
	return cd.Trie.Get(key).(dictionary.Attribute)
}

func (cd *CoreDictionary) GetByWordID(wordID int) dictionary.Attribute {
	return cd.Trie.Get3(wordID).(dictionary.Attribute)
}

func (cd *CoreDictionary) GetWordID(a string) int {
	return cd.Trie.ExactMatchSearch(a)
}




