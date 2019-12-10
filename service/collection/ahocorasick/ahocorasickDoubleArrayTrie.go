package ahocorasick

import (
	. "eagle/service/common"
	"eagle/service/corpus/io"
	"gifs/service/corpus/io/byteArray"
	"github.com/gogf/gf/container/glist"
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
func (a *AhoCorasickDoubleArrayTrie) ParseText(text string) glist.List {
	var position int = 1
	var currentState int = 0
	collectedEmits := glist.New(true)
	for i := 0; i < len(text); i++ {
		currentState = a.GetState(currentState, []Char(text)[i])
		a.StoreEmits(position, currentState, collectedEmits)
		position++
	}
	return *collectedEmits
}

// 保存输出
func (a *AhoCorasickDoubleArrayTrie) StoreEmits(position int, currentState int, collectedEmits *glist.List) {
	var hitArray []int = a.output[currentState]
	if hitArray != nil {
		for _, hit := range hitArray {
			collectedEmits.PushBack(NewHit(position-a.l[hit], position, a.v[hit]))
		}
	}
}

// 处理文本
func (a *AhoCorasickDoubleArrayTrie) ParseText2(text string, processor IHit) {
	var position int = 1
	var currentState int = 0
	textCharArray := String(text).ToCharArray()
	for i := 0; i < len(textCharArray); i++ {
		currentState = a.GetState(currentState, textCharArray[i])
		var hitArray []int = a.output[currentState]
		if hitArray != nil {
			for _, hit := range hitArray {
				processor.Hit(position-a.l[hit], position, a.v[hit])
			}
		}
		position++
	}
}

// 处理文本
func (a *AhoCorasickDoubleArrayTrie) ParseText3(text []Char, processor IHit) {
	var position int = 1
	var currentState int = 0
	for _, c := range text {
		currentState = a.GetState(currentState, c)
		var hitArray []int = a.output[currentState]
		if hitArray != nil {
			for _, hit := range hitArray {
				processor.Hit(position-a.l[hit], position, a.v[hit])
			}
		}
		position++
	}
}

/**
 * 持久化
 *
 * @param out 一个DataOutputStream
 * @throws Exception 可能的IO异常等
 */
func (a *AhoCorasickDoubleArrayTrie) Save(out io.DataOutputStream) {
	out.WriteInt(a.size)
	for i := 0; i < a.size; i++ {
		out.WriteInt(a.base[i])
		out.WriteInt(a.check[i])
		out.WriteInt(a.fail[i])
		var output []int = a.output[i]
		if output == nil {
			out.WriteInt(0)
		} else {
			out.WriteInt(len(output))
			for _, o := range output {
				out.WriteInt(o)
			}
		}
	}
	out.WriteInt(len(a.l))
	for _, length := range a.l {
		out.WriteInt(length)
	}
}

func (a *AhoCorasickDoubleArrayTrie) Save2() {

}

func (a *AhoCorasickDoubleArrayTrie) Load() {

}

/**
 * 载入
 *
 * @param byteArray 一个字节数组
 * @param value     值数组
 * @return 成功与否
 */
func (a *AhoCorasickDoubleArrayTrie) Load2(byteArray ByteArray.ByteArray, value []interface{}) bool {
	if byteArray.GetLength() == 0 {
		return false
	}
	a.size = byteArray.NextInt()
	a.base = make([]int, 0, a.size+65535)
	a.check = make([]int, 0, a.size+65535)
	a.fail = make([]int, 0, a.size+65535)
	a.output = make([][]int, 0, a.size+65535)
	var length int
	for i := 0; i < a.size; i++ {
		a.base[i] = byteArray.NextInt()
		a.check[i] = byteArray.NextInt()
		a.fail[i] = byteArray.NextInt()
		length = byteArray.NextInt()
		if length == 0 {
			continue
		}
		a.output[i] = make([]int, 0, length)
		for j := 0; j < len(a.output[i]); j++ {
			a.output[i][j] = byteArray.NextInt()
		}
	}
	length = byteArray.NextInt()
	a.l = make([]int, 0, length)
	for i := 0; i < len(a.l); i++ {
		a.l[i] = byteArray.NextInt()
	}
	a.v = value
	return true
}

/**
 * 获取值
 *
 * @param key 键
 * @return
 */
func (a *AhoCorasickDoubleArrayTrie) Get(key string) interface{} {
	index := a.ExactMatchSearch(key)
	if index >= 0 {
		return a.v[index]
	}
	return nil
}

/**
 * 更新某个键对应的值
 *
 * @param key   键
 * @param value 值
 * @return 是否成功（失败的原因是没有这个键）
 */
func (a *AhoCorasickDoubleArrayTrie) Set(key string, value interface{}) bool {
	index := a.ExactMatchSearch(key)
	if index >= 0 {
		a.v[index] = value
		return true
	}
	return false
}

/**
 * 从值数组中提取下标为index的值<br>
 * 注意为了效率，此处不进行参数校验
 *
 * @param index 下标
 * @return 值
 */
func (a *AhoCorasickDoubleArrayTrie) Get2(index int) interface{} {
	return a.v[index]
}

func (a *AhoCorasickDoubleArrayTrie) GetState(currentState int, character Char) int {
	var newCurrentState int = a.TransitionWithRoot(currentState, character)
	for newCurrentState == -1 {
		currentState = a.fail[currentState]
		newCurrentState = a.TransitionWithRoot(currentState, character)
	}
	return currentState
}

/**
 * 转移状态
 *
 * @param current
 * @param c
 * @return
 */
func (a *AhoCorasickDoubleArrayTrie) Transition(current int, c Char) int {
	var b int = current
	var p int

	p = b + c.ToInt() + 1
	if b == a.check[p] {
		b = a.base[p]
	} else {
		return -1
	}
	p = b
	return p
}

/**
 * c转移，如果是根节点则返回自己
 *
 * @param nodePos
 * @param c
 * @return
 */
func (a *AhoCorasickDoubleArrayTrie) TransitionWithRoot(nodePos int, c Char) int {
	var b int = a.base[nodePos]
	var p int

	p = b + c.ToInt() + 1
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

/**
 * 精确匹配
 *
 * @param key 键
 * @return 值的下标
 */
func (a *AhoCorasickDoubleArrayTrie) ExactMatchSearch(key string) int {
	return a.ExactMatchSearch2(key, 0, 0, 0)
}

/**
 * 精确匹配
 *
 * @param key
 * @param pos
 * @param len
 * @param nodePos
 * @return
 */
func (a *AhoCorasickDoubleArrayTrie) ExactMatchSearch2(key string, pos int, length int, nodePos int) int {
	if length <= 0 {
		length = String(key).Length()
	}
	if nodePos <= 0 {
		nodePos = 0
	}

	var result int = -1

	keyChars := String(key).ToCharArray()
	var b int = a.base[nodePos]
	var p int

	for i := 0; i < length; i++ {
		p = b + keyChars[i].ToInt() + 1
		if b == a.check[p] {
			b = a.base[p]
		} else {
			return result
		}
	}

	p = b
	var n int = a.base[p]
	if b == a.check[p] && n < 0 {
		result = -n - 1
	}
	return result
}

/**
 * 精确查询
 *
 * @param keyChars 键的char数组
 * @param pos      char数组的起始位置
 * @param len      键的长度
 * @param nodePos  开始查找的位置（本参数允许从非根节点查询）
 * @return 查到的节点代表的value ID，负数表示不存在
 */
func (a *AhoCorasickDoubleArrayTrie) ExactMatchSearch3(keyChars []Char, pos int, length int, nodePos int) int {
	var result int = -1
	var b int = a.base[nodePos]
	var p int

	for i := pos; i < length; i++ {
		p = b + keyChars[i].ToInt() + 1
		if b == a.check[p] {
			b = a.base[p]
		} else {
			return result
		}
	}

	p = b
	var n int = a.base[p]
	if b == a.check[p] && n < 0 {
		result = -n - 1
	}
	return result
}

/**
 * 大小，即包含多少个模式串
 *
 * @return
 */
func (a *AhoCorasickDoubleArrayTrie) Size() int {
	if a.v == nil {
		return 0
	} else {
		return len(a.v)
	}
}
