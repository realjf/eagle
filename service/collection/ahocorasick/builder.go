package ahocorasick

import (
	. "eagle/service/common"
	"eagle/utils"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/util/gconv"
	"math"
)

type Builder struct {
	rootState    *State // 根节点，仅仅用于构建过程
	used         []bool // 是否占用，仅仅用于构建
	allocSize    int    // 已分配在内存中的大小
	progress     int    // 一个控制增长速度的变量
	nextCheckPos int    // 下一个插入的位置将从此开始搜索
	keySize      int    // 键值对的大小

	check []int         // 双数组之check
	base  []int         // 双数组之base
	v     []interface{} // 保存value
	l     []int         // 每个key的长度
	size  int           // base和check 的大小

	fail   []int   // fail表
	output [][]int // 输出表
}

func (b *Builder) Build(ma gmap.TreeMap) {
	b.v = ma.Values()
	b.l = make([]int, len(b.v))
	keySet := ma.Keys()
	// 构建二分trie树
	b.addAllKeyword(keySet)
	// 在二分trie树的基础上构建双数组trie树
	b.buildDoubleArrayTrie(keySet)
	b.used = nil
	// 构建failure表并且合并output表
	b.constructFailureStates()
	b.rootState = nil
}

/**
 * 添加一个键
 *
 * @param keyword 键
 * @param index   值的下标
 */
func (b *Builder) addKeyword(keyword String, index int) {
	currentState := b.rootState
	for _, character := range keyword.ToCharArray() {
		currentState = currentState.AddState(character)
	}
	currentState.AddEmit(index)
	b.l[index] = keyword.Length()
}

func (b *Builder) addAllKeyword(keywordSet []interface{}) {
	i := 0
	for _, keyword := range keywordSet {
		word := keyword.(string)
		b.addKeyword(String(word), i)
		i++
	}
}

/**
 * 建立failure表
 */
func (b *Builder) constructFailureStates() {
	b.fail = make([]int, b.size+1)
	b.fail[1] = b.base[0]
	b.output = make([][]int, b.size+1)
	queue := gqueue.New()

	// 第一步，将深度为1的节点的failure设为根节点
	for _, depthOneState := range b.rootState.GetStates() {
		depthOneState.setFailure(b.rootState, b.fail)
		queue.Push(depthOneState)
		b.constructOutput(*depthOneState)
	}

	// 第二步，为深度 > 1 的节点建立failure表，这是一个bfs
	for queue.Size() != 0 {
		currentState := queue.Pop()
		for _, transition := range currentState.(*State).GetTransitions() {
			targetState := currentState.(*State).NextState2(transition)
			queue.Push(targetState)

			traceFailureState := currentState.(*State).Failure()
			for traceFailureState.NextState2(transition) == nil {
				traceFailureState = traceFailureState.Failure()
			}
			newFailureState := traceFailureState.NextState2(transition)
			targetState.setFailure(newFailureState, b.fail)
			targetState.AddEmit2(*newFailureState.Emit())
			b.constructOutput(*targetState)
		}
	}
}

/**
 * 建立output表
 */
func (b *Builder) constructOutput(target State) {
	emit := target.Emit()
	if emit == nil || emit.Size() == 0 {
		return
	}
	output := make([]int, emit.Size())
	i := 0
	emit.Iterator(func(v int) bool {
		output[i] = v
		i++
		return true
	})
	b.output[target.GetIndex()] = output
}

func (b *Builder) buildDoubleArrayTrie(keySet []interface{}) {
	b.progress = 0
	b.keySize = len(keySet)
	b.resize(65536 * 32)

	b.base[0] = 1
	b.nextCheckPos = 0

	root_node := b.rootState
	siblings := garray.New(true)
	b.fetch(*root_node, siblings)
	b.insert(*siblings)
}

/**
 * 拓展数组
 *
 * @param newSize
 * @return
 */
func (b *Builder) resize(newSize int) int {
	base2 := make([]int, newSize)
	check2 := make([]int, newSize)
	used2 := make([]bool, newSize)
	if b.allocSize > 0 {
		copy(base2, b.base)
		copy(check2, b.check)
		copy(used2, b.used)
	}

	b.base = base2
	b.check = check2
	b.used = used2
	b.allocSize = newSize
	return newSize
}

/**
 * 插入节点
 *
 * @param siblings 等待插入的兄弟节点
 * @return 插入位置
 */
func (b *Builder) insert(siblings garray.Array) int {
	begin := 0
	var pos int = gconv.Int(math.Max(float64(siblings.Get(0).(EntrySet).GetKey().(int)+1), float64(b.nextCheckPos)) - 1)
	nonzero_num := 0
	first := 0

	if b.allocSize <= pos {
		b.resize(pos + 1)
	}

outer:
	// 此循环体的目标是找出满足base[begin + a1...an]  == 0的n个空闲空间,a1...an是siblings中的n个节点
	for {
		pos++
		if b.allocSize <= pos {
			b.resize(pos + 1)
		}

		if b.check[pos] != 0 {
			nonzero_num++
			continue
		} else if first == 0 {
			b.nextCheckPos = pos
			first = 1
		}

		// 当前位置离第一个兄弟节点的距离
		begin = pos - siblings.Get(0).(EntrySet).GetKey().(int)
		if b.allocSize <= (begin + siblings.Get(siblings.Len()-1).(EntrySet).GetKey().(int)) {
			var l float64
			if 1.05 > 1.0*float64(b.keySize)/float64(b.progress+1) {
				l = 1.05
			} else {
				l = 1.0 * float64(b.keySize) / float64(b.progress+1)
			}
			b.resize(int(float64(b.allocSize) * l))
		}

		if b.used[begin] {
			continue
		}

		for i := 1; i < siblings.Len(); i++ {
			if b.check[begin+siblings.Get(i).(EntrySet).GetKey().(int)] != 0 {
				continue outer
			}
		}
		break
	}

	if 1.0*float64(nonzero_num)/float64(pos-b.nextCheckPos+1) >= 0.95 {
		b.nextCheckPos = pos
	}
	b.used[begin] = true

	if b.size > begin+siblings.Get(siblings.Len()-1).(EntrySet).GetKey().(int)+1 {
		b.size = b.size
	} else {
		b.size = begin + siblings.Get(siblings.Len()-1).(EntrySet).GetKey().(int) + 1
	}

	for _, entry := range siblings.Slice() {
		entrySet := entry.(EntrySet)
		b.check[begin+entrySet.GetKey().(int)] = begin
	}

	for _, entry := range siblings.Slice() {
		entrySet := entry.(EntrySet)
		new_siblings := garray.New(true)
		if b.fetch(entrySet.GetValue().(State), new_siblings) == 0 { // 一个词的终止且不为其他词的前缀，其实就是叶子节点
			b.base[begin+entrySet.GetKey().(int)] = (-entrySet.GetValue().(State).GetLargestValueId() - 1)
			b.progress++
		} else {
			h := b.insert(*new_siblings)
			b.base[begin+entrySet.GetKey().(int)] = h
		}
		entrySet.GetValue().(State).SetIndex(begin + entrySet.GetKey().(int))
	}

	return begin
}

/**
 * 释放空闲的内存
 */
func (b *Builder) loseWeight() {
	var nbase []int = make([]int, b.size+65535)
	copy(nbase, b.base)
	b.base = nbase

	var nCheck []int = make([]int, b.size+65535)
	copy(nCheck, b.check)
	b.check = nCheck
}

func (b *Builder) fetch(parent State, siblings *garray.Array) int {
	if parent.IsAcceptable() {
		fakeNode := NewState2(parent.GetDepth() + 1)
		fakeNode.AddEmit(parent.GetLargestValueId())
		siblings.PushLeft(EntrySet{0: fakeNode})
	}
	entrySets := parent.GetSuccess().Map()
	for _, entry := range utils.MapToMapEntrySet(entrySets) {
		siblings.PushLeft(EntrySet{entry.GetKey().(int) + 1: entry.GetValue()})
	}
	return siblings.Len()
}
