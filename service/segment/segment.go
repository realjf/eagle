package segment

import (
	"eagle/service/collection/trie"
	"eagle/service/collection/trie/bintrie"
	. "eagle/service/common"
	"eagle/service/config"
	"eagle/service/dictionary"
	"eagle/service/dictionary/customDictionary"
	"eagle/service/dictionary/other/charType"
	"eagle/service/segment/NShort/path"
	. "eagle/service/segment/common"
	"eagle/service/utility/sentencesUtil"
	"github.com/gogf/gf/container/glist"
)

type Segment struct {
	Config *Config // 分词器配置
}

func NewSegment() *Segment {
	return &Segment{
		Config: NewConfig(),
	}
}

func (s *Segment) AtomSegment(charArray []Char, start int, end int) glist.List {
	return *glist.New(true)
}

func (s *Segment) SimpleAtomSegment(charArray []Char, start int, end int) glist.List {
	return *glist.New(true)
}

func (s *Segment) QuickAtomSegment(charArray []Char, start int, end int) glist.List {
	atomNodeList := glist.New(true)
	offsetAtom := start
	var preType int = charType.GCharType.Get(charArray[offsetAtom]).ToInt()
	var curType int

	for offsetAtom++; offsetAtom < end; offsetAtom++ {
		curType = charType.GCharType.Get(charArray[offsetAtom]).ToInt()
		if curType != preType {
			// 浮点数识别
			if preType == charType.GCharType.CT_NUM.ToInt() && String("，,．.").IndexOf(charArray[offsetAtom].ToString()) != -1 {
				if offsetAtom + 1 <end {
					var nextType int = charType.GCharType.Get(charArray[offsetAtom+1]).ToInt()
					if nextType == charType.GCharType.CT_NUM.ToInt() {
						continue
					}
				}
			}
			atomNodeList.PushBack(*path.NewAtomNode(NewString().Append(charArray[start:offsetAtom]).ToString(), preType))
			start = offsetAtom
		}
		preType = curType
	}
	if offsetAtom == end {
		atomNodeList.PushBack(*path.NewAtomNode(NewString().Append(charArray[start:offsetAtom]).ToString(), preType))
	}

	return *atomNodeList
}

/**
 * 分词
 *
 * @param text 待分词文本
 * @return 单词列表
 */
func (s *Segment) Seg(text []Char) glist.List {
	// @todo 正规化字符
	if config.GConfig.Normalization {

	}

	return s.SegSentence(text)
}

/**
 * 分词<br>
 * 此方法是线程安全的
 *
 * @param text 待分词文本
 * @return 单词列表
 */
func (s *Segment) Seg2(text string) glist.List {
	charArray := String(text).ToCharArray()
	// @todo 正规化

	if s.Config.ThreadNumber > 1 && len(charArray) > 10000 {
		sentenceList := sentencesUtil.ToSentenceList3(charArray)
		sentenceArray := make([]string, sentenceList.Size())

		termListArray := make([]*glist.List, len(sentenceArray))
		per := len(sentenceArray) / s.Config.ThreadNumber
		threadArray := make([]*WorkThread, s.Config.ThreadNumber)
		for i :=0; i < s.Config.ThreadNumber -1; i++ {
			from := i * per
			threadArray[i] = NewWorkThread(s, sentenceArray, termListArray, from, from + per)
			threadArray[i].Run()
		}
		threadArray[s.Config.ThreadNumber-1] = NewWorkThread(s, sentenceArray, termListArray, (s.Config.ThreadNumber - 1) * per, len(sentenceArray))
		threadArray[s.Config.ThreadNumber-1].Run()

		termList := glist.New(true)
		if s.Config.Offset || s.Config.IndexMode > 0 { // 由于分割了句子，所以需要重新校正offset
			sentenceOffset := 0
			for i := 0; i < len(sentenceArray); i++ {
				for _, term := range termListArray[i].FrontAll() {
					termp := term.(Term)
					termp.Offset += sentenceOffset
					termList.PushBack(term)
				}
				sentenceOffset += len(sentenceArray[i])
			}
		}else{
			for _, list := range termListArray {
				termList.PushBacks(list.FrontAll())
			}
		}

		return *termList
	}
	return s.SegSentence(charArray)
}

type WorkThread struct {
	seg ISegment
	SentenceArray []string
	termListArray []*glist.List
	From int
	To int
}

func NewWorkThread(seg ISegment, sentenceArray []string, termListArray []*glist.List, from int, to int) *WorkThread {
	return &WorkThread{
		seg: seg,
		SentenceArray: sentenceArray,
		termListArray: termListArray,
		From:          from,
		To:            to,
	}
}

func (wt *WorkThread) Run() {
	go func() {
		for i := wt.From; i < wt.To; i++ {
			result := wt.seg.SegSentence(String(wt.SentenceArray[i]).ToCharArray())
			wt.termListArray[i] = &result
		}
	}()
}

func (s *Segment) SegSentence(sentence []Char) glist.List {
	return *glist.New(true)
}

// 将一条路径转为最终结果
func (s *Segment) Convert(vertexList glist.List, offsetEnabled bool) glist.List {
	length := vertexList.Size() - 2
	resultList := glist.New(true)
	if offsetEnabled {
		offset := 0
		for i := 0; i <length; i++ {
			resultList.Iterator(func(e *glist.Element) bool {
				var vertex Vertex = e.Value.(Vertex)
				var term Term = s.ConvertFromVertex(vertex)
				term.Offset = offset
				offset += term.Length()
				resultList.PushBack(term)
				return true
			})
		}
	}else{
		for i := 0; i < length; i++ {
			resultList.Iterator(func(e *glist.Element) bool {
				var vertex Vertex = e.Value.(Vertex)
				var term Term = s.ConvertFromVertex(vertex)
				resultList.PushBack(term)
				return true
			})
		}
	}

	return *resultList
}

// 将节点转为term
func (s *Segment) ConvertFromVertex(vertex Vertex) Term {
	return *NewTerm(vertex.RealWord, vertex.GuessNature())
}

func (s *Segment) CombineByCustomDictionary(vertexList glist.List) glist.List {
	return s.CombineByCustomDictionary2(vertexList, customDictionary.GCustomDictionary.Dat)
}

func (s *Segment) CombineByCustomDictionary2(vertexList glist.List, dat *trie.DoubleArrayTrie) glist.List {
	wordNet := make([]Vertex, vertexList.Size())
	// DAT合并
	var length int = len(wordNet) - 1 // 跳过首尾
	for i := 1; i < length; i++ {
		var state int = 1
		state = dat.Transition3(wordNet[i].RealWord, state)
		if state > 0 {
			var to int = i+1
			end := to
			var value dictionary.Attribute = dat.Output(state).(dictionary.Attribute)
			for ; to < length; to++ {
				state = dat.Transition3(wordNet[i].RealWord, state)
				if state < 0 {
					break
				}
				output := dat.Output(state).(dictionary.Attribute)
				if &output != nil {
					value = output
					end = to + 1
				}
			}
			if &value != nil {
				s.CombineWords(wordNet, i, end, value)
				i = end -1
			}
		}
	}
	// BinTrie合并
	if customDictionary.GCustomDictionary.Trie != nil {
		for i := 1; i < length; i++ {
			if &wordNet[i] == nil {
				continue
			}
			var state bintrie.BaseNode = customDictionary.GCustomDictionary.Trie.
			if state != nil {
				var to int = i + 1
				var end int = to
				var value dictionary.Attribute = 

				for ; to < length; to++ {
					if &wordNet[to] == nil {
						continue
					}
					state = state
					if state == nil {
						break
					}
				}
				if value != nil {
					s.CombineWords(wordNet, i, end, value)
					i = end - 1
				}
			}
		}
	}

	vertexList.Clear()
	for _, vertex := range wordNet {
		if &vertex != nil {
			vertexList.PushBack(vertex)
		}
	}
	return vertexList
}

func (s *Segment) CombineByCustomDictionary3(vertexList glist.List, wordNetAll WordNet) glist.List {
	return s.CombineByCustomDictionary4(vertexList, customDictionary.GCustomDictionary.Dat, wordNetAll)
}

/**
 * 使用用户词典合并粗分结果，并将用户词语收集到全词图中
 * @param vertexList 粗分结果
 * @param dat 用户自定义词典
 * @param wordNetAll 收集用户词语到全词图中
 * @return 合并后的结果
 */
func (s *Segment) CombineByCustomDictionary4(vertexList glist.List, dat *trie.DoubleArrayTrie, wordNetAll WordNet) glist.List {
	outputList := s.CombineByCustomDictionary2(vertexList, dat)
	line := 0
	for _, vertex := range outputList.FrontAll() {
		parentLength := String(vertex.(Vertex).RealWord).Length()
		currentLine := line
		if parentLength >= 3 {

		}
		line += parentLength
	}
	return outputList
}

/**
 * 将连续的词语合并为一个
 * @param wordNet 词图
 * @param start 起始下标（包含）
 * @param end 结束下标（不包含）
 * @param value 新的属性
 */
func (s *Segment) CombineWords(wordNet []Vertex, start int, end int, value dictionary.Attribute) {
	if start + 1 == end {
		wordNet[start].Attribute = value
	}else{
		sbTerm := NewString()
		for j := start; j < end ; j++ {
			if &wordNet[j] == nil {
				continue
			}
			realWord := wordNet[j].RealWord
			sbTerm.Append(realWord)
			wordNet[j] = Vertex{}
		}
		realWord := sbTerm.ToString()
		wordNet[start] = *NewVertex(realWord, realWord, value)
	}
}

