package MDAG

import (
	"fmt"
	. "eagle/service/common"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
	"github.com/emirpasic/gods/stacks/arraystack"
)

/**
 * MDAG中的一个节点<br>
 */
type MDAGNode struct {
	isAcceptNode bool // 是否是终止状态
	outgoingTransitionTreeMap *gmap.TreeMap // 状态转移函数
	incomingTransitionCount int // 入度
	transitionSetBeginIndex int // 在简化的MDAG中表示该节点的转移状态集合的起始位置
	storedHashCode int // 当它被计算后的hash值
}

// 建立一个节点
// isAcceptNode     是否是终止状态
func NewMDAGNode(isAcceptNode bool) *MDAGNode {
	mdagn := &MDAGNode{
		isAcceptNode: isAcceptNode,
		transitionSetBeginIndex: -1,
		incomingTransitionCount: 0,
		storedHashCode: 0,
		outgoingTransitionTreeMap: gmap.NewTreeMap(gutil.ComparatorByte, true),
	}

	return mdagn
}

// 克隆一个状态
func NewMDAGNodeFromNode(node MDAGNode) *MDAGNode {
	mdagn := &MDAGNode{
		isAcceptNode:              node.isAcceptNode,
		outgoingTransitionTreeMap: gmap.NewTreeMap(gutil.ComparatorByte, true),
		incomingTransitionCount:   0,
		transitionSetBeginIndex:   -1,
		storedHashCode:            0,
	}

	for key, transitionKeyValuePair := range mdagn.outgoingTransitionTreeMap.Map() {
		node := transitionKeyValuePair.(MDAGNode)
		node.incomingTransitionCount++
		mdagn.outgoingTransitionTreeMap.Set(key, node)
	}
	return mdagn
}

// 克隆一个状态
func (m *MDAGNode) Clone() MDAGNode {
	return *NewMDAGNodeFromNode(*m)
}

// 克隆一个状态
func (m *MDAGNode) Clone2(soleParentNode MDAGNode, parentToCloneTransitionLabelChar Character) MDAGNode {
	cloneNode := NewMDAGNodeFromNode(*m)
	soleParentNode.ReassignOutgoingTransition(parentToCloneTransitionLabelChar, m, *cloneNode)
	return *cloneNode
}

func (m *MDAGNode) GetTransitionSetBeginIndex() int {
	return m.transitionSetBeginIndex
}

func (m *MDAGNode) GetOutgoingTransitionCount() int {
	return m.outgoingTransitionTreeMap.Size()
}

func (m *MDAGNode) GetIncomingTransitionCount() int {
	return m.incomingTransitionCount
}

func (m *MDAGNode) IsConfluenceNode() bool {
	return m.incomingTransitionCount > 1
}

func (m *MDAGNode) IsAcceptNode() bool {
	return m.isAcceptNode
}

func (m *MDAGNode) SetAcceptStateStatus(isAcceptNode bool) {
	m.isAcceptNode = isAcceptNode
}

// 转移状态在数组中的起始下标
func (m *MDAGNode) SetTransitionSetBeginIndex(transitionSetBeginIndex int) {
	m.transitionSetBeginIndex = transitionSetBeginIndex
}

func (m *MDAGNode) HasOutgoingTransition(letter Character) bool {
	return m.outgoingTransitionTreeMap.Contains(letter)
}

func (m *MDAGNode) HasTransitions() bool {
	return !m.outgoingTransitionTreeMap.IsEmpty()
}

func (m *MDAGNode) Transition(letter Character) MDAGNode {
	return m.outgoingTransitionTreeMap.Get(letter).(MDAGNode)
}

func (m *MDAGNode) TransitionFromStr(str string) MDAGNode {
	charCount := len(str)
	currentNode := *m

	for i:=0; i<charCount; i++ {
		currentNode = currentNode.Transition(String(str).ToCharArray()[i])
		if &currentNode == nil {
			break
		}
	}

	return currentNode
}

func (m *MDAGNode) TransitionFromCharArray(str []Character) MDAGNode {
	charCount := len(str)
	currentNode := *m
	for i:=0; i < charCount; i++ {
		currentNode = currentNode.Transition(str[i])
		if &currentNode == nil {
			break
		}
	}

	return currentNode
}

func (m *MDAGNode) TransitionFromCharArrayOffset(str []Character, offset int) MDAGNode {
	charCount := len(str)
	currentNode := *m
	for i:=0; i < charCount; i++ {
		currentNode = currentNode.Transition(str[i+offset])
		if &currentNode == nil {
			break
		}
	}

	return currentNode
}

// 获取一个字符串路径上经过的节点
func (m *MDAGNode) GetTransitionPathNodes(str string) arraystack.Stack {
	nodeStack := arraystack.New()
	currentNode := *m
	numberOfChars := len(str)

	for i := 0; i<numberOfChars && &currentNode != nil; i++ {
		currentNode = currentNode.Transition(String(str).ToCharArray()[i])
		nodeStack.Push(currentNode)
	}

	return *nodeStack
}

func (m *MDAGNode) GetOutgoingTransitions() gmap.TreeMap {
	return *m.outgoingTransitionTreeMap
}

// 本状态的目标状态们的入度减一
func (m *MDAGNode) DecrementTargetIncomingTransitionCounts() {
	for key, transitionKeyValuePair := range m.outgoingTransitionTreeMap.Map() {
		node := transitionKeyValuePair.(MDAGNode)
		node.incomingTransitionCount--
		m.outgoingTransitionTreeMap.Set(key, node)
	}
}

// 重新设置转移状态函数的目标
func (m *MDAGNode) ReassignOutgoingTransition(letter Character, oldTargetNode *MDAGNode, newTargetNode MDAGNode) {
	oldTargetNode.incomingTransitionCount--
	newTargetNode.incomingTransitionCount++

	m.outgoingTransitionTreeMap.Set(letter, newTargetNode)
}

// 新建一个转移目标
func (m *MDAGNode) AddOutgoingTransition(letter Character, targetAcceptStateStatus bool) MDAGNode {
	newTargetNode := NewMDAGNode(targetAcceptStateStatus)
	newTargetNode.incomingTransitionCount++

	m.outgoingTransitionTreeMap.Set(letter, *newTargetNode)
	return *newTargetNode
}

/**
 * 建立一条边（起点是自己）
 * @param letter 边上的字符串
 * @param newTargetNode 边的重点
 * @return 终点
 */
func (m *MDAGNode) AddOutgoingTransitionFromNode(letter Character, newTargetNode MDAGNode) MDAGNode {
	newTargetNode.incomingTransitionCount++
	m.outgoingTransitionTreeMap.Set(letter, newTargetNode)
	return newTargetNode
}

// 移除一个转移目标
func (m *MDAGNode) RemoveOutgoingTransition(letter Character) {
	m.outgoingTransitionTreeMap.Remove(letter)
}

// 是否含有相同的转移函数
func HaveSameTransitions(node1 MDAGNode, node2 MDAGNode) bool {
	outgoingTransitionTreeMap1 := node1.outgoingTransitionTreeMap
	outgoingTransitionTreeMap2 := node2.outgoingTransitionTreeMap

	if outgoingTransitionTreeMap1.Size() == outgoingTransitionTreeMap2.Size() {
		for key, transitionKeyValuePair := range outgoingTransitionTreeMap1.Map() {
			currentTargetNode := transitionKeyValuePair.(MDAGNode)
			currentCharKey := key.(Character)
			if !outgoingTransitionTreeMap2.Contains(currentCharKey) || !outgoingTransitionTreeMap2.Get(currentCharKey).(MDAGNode).Equals(currentTargetNode) {
				return false
			}
		}
	}else{
		return false
	}

	return true
}

func (m *MDAGNode) ClearStoredHashCode() {
	m.storedHashCode = 0
}

// 两个状态是否等价，只有状态转移函数完全一致才算相等
func (m *MDAGNode) Equals(v interface{}) bool {
	areEqual := *m == v.(MDAGNode)
	if !areEqual && v != nil {
		if node, ok := v.(MDAGNode); ok {
			if node.isAcceptNode == m.isAcceptNode && HaveSameTransitions(*m, node) {
				areEqual = true
			}else{
				areEqual = false
			}
		}else{
			areEqual = false
		}
	}

	return areEqual
}

func (m *MDAGNode) HashCode() int {
	if m.storedHashCode == 0 {
		hash := 7
		if m.isAcceptNode {
			hash = 53 * hash + 1
		}
		if m.outgoingTransitionTreeMap != nil {
			hash = 53 * hash + GetTreeMapHashCode(*m.outgoingTransitionTreeMap)
		}

		m.storedHashCode = hash
		return hash
	}else{
		return m.storedHashCode
	}
}

func (m *MDAGNode) ToString() string {
	return fmt.Sprintf("MDAGNode{isAcceptNode=%v, outgoingTransitionTreeMap=%v, incomingTransitionCount=%d}",
		m.isAcceptNode, m.outgoingTransitionTreeMap.Keys(), m.incomingTransitionCount)
}

func GetTreeMapHashCode(treeMap gmap.TreeMap) int {
	h := 0
	for key, value := range treeMap.Map() {
		keyHash, valueHash := 0, 0
		if key != nil {
			keyHash = key.(Character).ToInt()
		}
		if value != nil {
			valueHash = value.(MDAGNode).HashCode()
		}
		h += keyHash ^ valueHash
	}
	return h
}

