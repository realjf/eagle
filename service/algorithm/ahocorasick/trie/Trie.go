package trie

import (
	"fmt"
	"eagle/service/algorithm/ahocorasick/interval"
	. "eagle/service/common"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/text/gstr"
)

// 基于 Aho-Corasick 白皮书, 贝尔实验室: ftp://163.13.200.222/assistant/bearhero/prog/%A8%E4%A5%A6/ac_bm.pdf
type Trie struct {
	trieConfig TrieConfig
	rootState  State

	failureStatesConstructed bool // 是否建立了failure表
}

// 构造一棵trie树
func NewTrie(trieConfig TrieConfig) *Trie {
	return &Trie{
		trieConfig: trieConfig,
		rootState:  *NewState(),
	}
}

func NewTrie2() *Trie {
	return &Trie{}
}

func NewTrie3(keywords []string) *Trie {
	trie := Trie{}
	trie.AddAllKeyword(keywords)
	return &trie
}

func (t *Trie) RemoveOverlaps() *Trie {
	t.trieConfig.SetAllowOverlaps(false)
	return t
}

// 只保留最长匹配
func (t *Trie) RemainLongestEnable() *Trie {
	t.trieConfig.RemainLongest = true
	return t
}

func (t *Trie) AddKeyword(keyword string) {
	if keyword == "" || len(keyword) == 0 {
		return
	}
	currentState := t.rootState
	// for
	currentState.AddEmit(keyword)
}

func (t *Trie) AddAllKeyword(keywordSet []string) {
	for _, keyword := range keywordSet {
		t.AddKeyword(keyword)
	}
}

// 一个最长分词器
// text 待分词文本
func (t *Trie) Tokenize(text string) glist.List {
	tokens := glist.New(true)
	collectedEmits := t.ParseText(text)
	// 下面是最长分词的关键
	intervalTree := interval.NewIntervalTree(collectedEmits)
	intervalTree.RemoveOverlaps(collectedEmits)
	// 移除结束

	lastCollectedPosition := -1
	for _, iemit := range collectedEmits.FrontAll() {
		emit := iemit.(Emit)
		if emit.GetStart()-lastCollectedPosition > 1 {
			tokens.PushBack(t.createFragment(emit, text, lastCollectedPosition))
		}
		tokens.PushBack(t.createMatch(emit, text))
		lastCollectedPosition = emit.GetEnd()
	}
	if len(text)-lastCollectedPosition > 1 {
		tokens.PushBack(t.createFragment(Emit{}, text, lastCollectedPosition))
	}

	return *tokens
}

func (t *Trie) createFragment(emit Emit, text string, lastCollectedPosition int) IToken {
	start := 0
	if &emit == nil {
		start = len(text)
	} else {
		start = emit.GetStart()
	}
	fragment := gstr.SubStr(text, lastCollectedPosition+1, start)
	return NewFragmentToken(fragment)
}

func (t *Trie) createMatch(emit Emit, text string) IToken {
	str := gstr.SubStr(text, emit.GetStart(), emit.GetEnd()+1)
	return NewMatchToken(str, emit)
}

// 模式匹配
// text 待匹配的文本
func (t *Trie) ParseText(text string) glist.List {
	t.checkForConstructedFailureStates()

	position := 0
	currentState := t.rootState
	collectedEmits := glist.New()
	for i := 0; i < len(text); i++ {
		character := String(text).ToCharArray()[i]
		currentState = t.GetState(currentState, character)
		t.StoreEmits(position, currentState, *collectedEmits)
		position++
	}

	if !t.trieConfig.IsAllowOverlaps() {
		intervalTree := interval.NewIntervalTree(*collectedEmits)
		intervalTree.RemoveOverlaps(*collectedEmits)
	}

	if t.trieConfig.RemainLongest {
		t.RemainLongest(*collectedEmits)
	}

	return *collectedEmits
}

// 只保留最长词
func (t *Trie) RemainLongest(collectedEmits glist.List) {
	if collectedEmits.Size() < 2 {
		return
	}
	emitMapStart := gmap.TreeMap{}
	for _, iemit := range collectedEmits.FrontAll() {
		emit := iemit.(Emit)
		pre := emitMapStart.Get(emit.GetStart()).(Emit)
		if &pre == nil || pre.Size() < emit.Size() {
			emitMapStart.Set(emit.GetStart(), emit)
		}
	}
	if emitMapStart.Size() < 2 {
		collectedEmits.Clear()
		collectedEmits.PushBacks(emitMapStart.Values())
		return
	}
	emitMapEnd := gmap.TreeMap{}
	for _, iemit := range emitMapStart.Values() {
		emit := iemit.(Emit)
		pre := emitMapEnd.Get(emit.GetEnd()).(Emit)
		if &pre == nil || pre.Size() < emit.Size() {
			emitMapEnd.Set(emit.GetEnd(), emit)
		}
	}

	collectedEmits.Clear()
	collectedEmits.PushBacks(emitMapEnd.Values())
}

// 跳转到下一个状态
// currentState 当前状态
// character    接受字符
func (t *Trie) GetState(currentState State, character Char) State {
	newCurrentState := currentState.NextStateIgnoreRootState(character)
	for &newCurrentState == nil {
		currentState = currentState.Failure()
		newCurrentState = currentState.NextState2(character)
	}
	return newCurrentState
}

// 检查是否建立了failure表
func (t *Trie) checkForConstructedFailureStates() {
	if !t.failureStatesConstructed {
		t.constructFailureStates()
	}
}

// 建立failure表
func (t *Trie) constructFailureStates() {
	queue := gqueue.New()

	// 第一步，将深度为1的节点的failure设为根节点
	for _, depthOneState := range t.rootState.GetStates() {
		depthOneState.SetFailure(t.rootState)
		queue.Push(depthOneState)
	}
	t.failureStatesConstructed = true

	// 第二步，为深度 > 1 的节点建立failure表，这是一个bfs
	for queue.Size() != 0 {
		currentState := queue.Pop().(State)

		for _, transition := range currentState.GetTransitions() {
			targetState := currentState.NextState2(transition)
			queue.Push(targetState)

			traceFailureState := currentState.Failure()
			for traceFailureState.NextState2(transition) != (State{}) {
				traceFailureState = traceFailureState.Failure()
			}

			newFailureState := traceFailureState.NextState2(transition)
			targetState.SetFailure(newFailureState)
			targetState.AddEmit2(newFailureState.Emit().FrontAll())
		}
	}
}

type IWalker interface {
	Meet(path string, state State) // 遇到了一个节点
}

func (t *Trie) Dfs(walker IWalker) {
	t.checkForConstructedFailureStates()
	t.Dfs2(t.rootState, "", walker)
}

func (t *Trie) Dfs2(currentState State, path string, walker IWalker) {
	walker.Meet(path, currentState)
	for _, transition := range currentState.GetTransitions() {
		targetState := currentState.NextState2(transition)
		t.Dfs2(targetState, fmt.Sprintf("%s%v", path, transition), walker)
	}
}

// 保存匹配结果
// position       当前位置，也就是匹配到的模式串的结束位置+1
// currentState   当前状态
// collectedEmits 保存位置
func (t *Trie) StoreEmits(position int, currentState State, collectedEmits glist.List) {
	emits := currentState.Emit()
	if (emits != glist.List{}) && emits.Len() != 0 {
		for _, iemit := range emits.FrontAll() {
			emit := iemit.(string)
			collectedEmits.PushBack(NewEmit(position-len(emit)+1, position, emit))
		}
	}
}

// 文本是否包含任何模式
// text 待匹配的文本
func (t *Trie) HasKeyword(text string) bool {
	t.checkForConstructedFailureStates()

	currentState := t.rootState
	for i := 0; i < len(text); i++ {
		character := String(text).ToCharArray()[i]
		nextState := t.GetState(currentState, character)
		if (nextState != State{}) && nextState != currentState && nextState.Emit().Size() != 0 {
			return true
		}
		currentState = nextState
	}

	return false
}
