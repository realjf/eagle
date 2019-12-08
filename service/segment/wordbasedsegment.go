package segment

import (
	"eagle/service/dictionary"
	"eagle/service/dictionary/coreDictionary"
	. "eagle/service/segment/common"
)

type WordBasedSegment struct {
	Segment
}

// 生成一元词网
func (wbs *WordBasedSegment) GenerateWordNet(wordNetStorage *WordNet) {
	charArray := wordNetStorage.CharArray
	// 核心词典查询
	searcher := coreDictionary.GCoreDictionary.Trie.GetSearcher3(charArray, 0)
	for searcher.Next() {
		wordNetStorage.Add(
			searcher.Begin+1,
			*NewVertex4(
				string(charArray[searcher.Begin:searcher.Length]),
				searcher.Value.(dictionary.Attribute),
				searcher.Index))
	}
	// 强制用户词典查询

	// 原子分词，保证图连通
	vertexes := wordNetStorage.GetVertexes()
	for i := 1; i < len(vertexes); {
		if vertexes[i].Size() == 0 {
			j := i+ 1
			for ; j < len(vertexes) -1; j++ {
				if vertexes[j].Size() != 0 {
					break
				}
			}
			wordNetStorage.AddSegment(i, wbs.QuickAtomSegment(charArray, i-1, j-1))
			i = j
		}else{
			i += len(vertexes[i].Back().Value.(Vertex).RealWord)
		}
	}
}
