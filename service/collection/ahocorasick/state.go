package ahocorasick

import (
	"fmt"
	. "eagle/service/common"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/container/gtree"
)

type SuccessTree struct {
	success gmap.TreeMap
	list *gtree.AVLTree
}

// 一个状态有如下几个功能
// success: 成功转移到另一个状态
// failure: 不可顺着字符串跳转的话，则跳转到一个浅一点的节点
// emits: 命中一个模式串
//
// 根节点稍有不同，根节点没有failure功能，它的failure指的是按照字符串路径转移到下一个状态，其他节点则都有failure状态
type State struct {
	depth int // 模式串的长度，也是这个状态的深度
	failure *State // fail函数，如果没有匹配到，则跳转到此状态
	emits *gset.IntSet // 只要这个状态可达，则记录模式串
	success *gmap.TreeMap // goto 表，也称转移函数。根据字符串的下一个字符转移到下一个状态

	index int // 在双数组中的对应下标
}

func NewState() *State {
	return &State{
		depth: 0,
	}
}

// 获取节点深度
func (s *State) GetDepth() int {
	return s.depth
}

// 添加一个匹配到的模式串（这个状态对应这这个模式串）
func (s *State) AddEmit(keyword int) {
	if s.emits == nil {
		s.emits = gset.NewIntSet(true)
	}
	s.emits.Add(keyword)
}

// 获取最大的值
func (s *State) GetLargestValueId() int {
	if s.emits == nil || s.emits.Size() == 0 {
		return 0
	}

	max := 0
	s.emits.Iterator(func(v int) bool {
		if v > max {
			max = v
		}
		return true
	})

	return max
}

// 是否是终止状态
func (s *State) IsAcceptable() bool {
	return s.depth > 0 && s.emits != nil
}

// 获取failure状态
func (s *State) Failure() *State {
	return s.failure
}

// 设置failure状态
func (s *State) setFailure(failState *State, fail []int) {
	s.failure = failState
	fail[s.index] = failState.index
}

// 转移到下一个状态
// character 希望按此字符转移
// ignoreRootState 是否忽略根节点，如果是根节点自己调用则应该是true，否则为false
func (s *State) NextState(character Char, ignoreRootState bool) *State {
	nextState := s.success.Get(character).(*State)
	if !ignoreRootState && nextState == nil && s.depth == 0 {
		nextState = s
	}

	return nextState
}

// 按照character转移，根节点转移失败会返回自己（永远不会返回null）
func (s *State) NextState2(character Char) *State {
	return s.NextState(character, false)
}

// 按照character转移，任何节点转移失败会返回null
func (s *State) NextStateIgnoreRootState(character Char) *State {
	return s.NextState(character, true)
}

func (s *State) AddState(character Char) *State {
	nextState := s.NextStateIgnoreRootState(character)
	if nextState == nil {
		nextState = &State{depth:s.depth+1}
		s.success.Set(character, nextState)
	}
	return nextState
}

func (s *State) GetStates() []*State {
	values := s.success.Values()

	newValues := []*State{}
	for _, v := range values {
		newValues = append(newValues, v.(*State))
	}

	return newValues
}

func (s *State) GetTransitions() []Char {
	keys := s.success.Keys()
	newKeys := []Char{}
	for _, v := range keys {
		newKeys = append(newKeys, v.(Char))
	}

	return newKeys
}

func (s *State) ToString() string {

	return fmt.Sprintf("State{depth=%d, ID=%d, emits=%v, success=%v, failureID=%d, failure=%v}",
		s.depth, s.index, s.emits, s.success, s.failure.index, s.failure)
}

// 获取goto表
func (s *State) GetSuccess() gmap.TreeMap {
	return *s.success
}

func (s *State) SetIndex(index int) {
	s.index = index
}

func (s *State) GetIndex() int {
	return s.index
}
