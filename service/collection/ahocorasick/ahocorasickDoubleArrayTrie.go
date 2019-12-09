package ahocorasick

import (
	. "eagle/service/common"
	"github.com/gogf/gf/container/gmap"
)

type AhoCorasickDoubleArrayTrie struct {
	Builder
}

func NewAhoCorasickDoubleArrayTrie() *AhoCorasickDoubleArrayTrie {
	return &AhoCorasickDoubleArrayTrie{}
}

func NewAhoCorasickDoubleArrayTrie2(dictionary gmap.TreeMap) *AhoCorasickDoubleArrayTrie {
	dat := NewAhoCorasickDoubleArrayTrie()
	dat.Build(dictionary)
	return dat
}

/**
 * 匹配母文本
 *
 * @param text 一些文本
 * @return 一个pair列表
 */
func (a *AhoCorasickDoubleArrayTrie) ParseText(text string) []Hit {
	var position int = 1
	var currentState int = 0
	collectedEmits := []Hit{}
	for i := 0; i < len(text); i++ {
		currentState = a.GetState(currentState, []Char(text)[i])
		a.StoreEmits(position, currentState, collectedEmits)
		position++
	}
	return collectedEmits
}

func (a *AhoCorasickDoubleArrayTrie) StoreEmits(position int, currentState int, collectedEmits []Hit) {
	var hitArray []int = a.output[currentState]
	if hitArray != nil {
		for _, hit := range hitArray {
			collectedEmits
		}
	}
}

func (a *AhoCorasickDoubleArrayTrie) ParseText2(text string, processor IHit) {
	var position int = 1
	var currentState int = 0
	for i := 0; i < len(text); i++ {
		currentState = a.GetState(currentState, []Char(text)[i])
		var hitArray []int = a.output[currentState]
		if hitArray != nil {
			for _, hit := range hitArray {
				processor.Hit(position - a.l[hit], position, a.v[hit])
			}
		}
		position++
	}
}

func (a *AhoCorasickDoubleArrayTrie) GetState(currentState int, character Char) int {
	var newCurrentState int = a.TransitionWithRoot(currentState, character)
	for ;newCurrentState == -1; {
		currentState = a.fail[currentState]
		newCurrentState = a.TransitionWithRoot(currentState, character)
	}
	return currentState
}

func (a *AhoCorasickDoubleArrayTrie) TransitionWithRoot(nodePos int, c Char) int {
	var b int = a.base[nodePos]
	var p int
	d := c
	p = b +  + 1
	if b != a.check[p] {
		if nodePos == 0 {
			return 0
		}
		return -1
	}
	return p
}

func (a *AhoCorasickDoubleArrayTrie) Build(m gmap.TreeMap) {
	a.Builder.Build(m)
}


func (a *AhoCorasickDoubleArrayTrie) constructOutput(targetState State) {

}

// 扩展数组
func (a *AhoCorasickDoubleArrayTrie) resize(newSize int) int {
	var base2 []int = make([]int, newSize)
	var check2 []int = make([]int, newSize)
	var used2 []bool = make([]bool, newSize)
	if a.allocSize > 0 {
		copy(base2, a.base)
		copy(check2, a.check)
		copy(used2, a.used)
	}

	a.base = base2
	a.check = check2
	a.used = used2

	a.allocSize = newSize
	return a.allocSize
}

// 释放空闲的内存
func (a *AhoCorasickDoubleArrayTrie) loseWeight() {
	var nbase []int = make([]int, a.size + 65535)
	copy(nbase, a.base)
	a.base = nbase

	var ncheck []int = make([]int, a.size + 65535)
	copy(ncheck, a.check)
	a.check = ncheck
}


