package trie

import (
	"fmt"
	. "eagle/service/common"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gconv"
)

type State struct {
	depth   int           // 模式串的长度，也是这个状态的深度
	failure *State        // fail 函数，如果没有匹配到，则跳转到此状态。
	emits   *glist.List   // 只要这个状态可达，则记录模式串
	success *gmap.TreeMap // goto 表，也称转移函数。根据字符串的下一个字符转移到下一个状态
}

// 构造深度为0的节点
func NewState() *State {
	return &State{
		depth:   0,
		failure: nil,
		emits:   nil,
		success: nil,
	}
}

// 构造深度为depth的节点
func NewState2(depth int) *State {
	return &State{
		depth:   depth,
		failure: nil,
		emits:   nil,
		success: nil,
	}
}

// 获取节点深度
func (s *State) GetDepth() int {
	return s.depth
}

// 添加一个匹配到的模式串（这个状态对应着这个模式串)
func (s *State) AddEmit(keyword string) {
	if s.emits == nil {
		s.emits = glist.New(true)
	}
	s.emits.PushBack(keyword)
}

// 添加一些匹配到的模式串
func (s *State) AddEmit2(emits []interface{}) {
	for _, emit := range emits {
		s.AddEmit(gconv.String(emit))
	}
}

func (s *State) Emit() glist.List {
	if s.emits == nil {
		return glist.List{}
	}
	return *s.emits
}

// 获取failure状态
func (s *State) Failure() State {
	return *s.failure
}

// 设置failure状态
func (s *State) SetFailure(failState State) {
	s.failure = &failState
}

// 转移到下一个状态
// character 希望按此字符转移
// ignoreRootState 是否忽略根节点，如果是根节点自己调用则应该是true，否则为false
func (s *State) NextState(character Char, ignoreRootState bool) State {
	nextState := s.success.Get(character).(*State)
	if !ignoreRootState && nextState == nil && s.depth == 0 {
		nextState = s
	}
	return *nextState
}

// 按照character转移，根节点转移失败会返回自己（永远不会返回null）
func (s *State) NextState2(character Char) State {
	return s.NextState(character, false)
}

// 按照character转移，任何节点转移失败会返回nil
func (s *State) NextStateIgnoreRootState(character Char) State {
	return s.NextState(character, true)
}

func (s *State) AddState(character Char) State {
	nextState := s.NextStateIgnoreRootState(character)
	if &nextState == nil {
		nextState = *NewState2(s.depth + 1)
		s.success.Set(character, nextState)
	}
	return nextState
}

func (s *State) GetStates() []State {
	result := []State{}
	for _, v := range s.success.Values() {
		result = append(result, v.(State))
	}
	return result
}

func (s *State) GetTransitions() []Char {
	result := []Char{}
	for _, v := range s.success.Keys() {
		result = append(result, v.(Char))
	}
	return result
}

func (s *State) ToString() string {
	return fmt.Sprintf("State{depth=%d, emits=%s, success=%s, failure=%v}",
		s.depth, s.emits.String(), s.success.Keys(), s.failure)
}
