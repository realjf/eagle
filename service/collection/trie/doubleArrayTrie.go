package trie

import (
	"bytes"
	"fmt"
	"eagle/service/collection/ahocorasick"
	. "eagle/service/common"
	"eagle/service/corpus/io"
	"eagle/service/corpus/io/byteArrayStream"
	"eagle/service/entry"
	"eagle/utils"
	"eagle/service/utility/byteUtility"
	ByteArrayStream "eagle/service/corpus/io/byteArrayStream"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gconv"
	"github.com/willf/bitset"
	"math"
	"os"
)

const (
	BUF_SIZE  = 16384
	UNIT_SIZE = 8
)

type Node struct {
	Code  int
	Depth int
	Left  int
	Right int
}

func (n *Node) ToString() string {
	return fmt.Sprintf("Node{code=%d, depth=%d, left=%d, right=%d}", n.Code, n.Depth, n.Left, n.Right)
}

type DoubleArrayTrie struct {
	BUF_SIZE  int
	UNIT_SIZE int

	// base 和 check 的大小
	size      int
	allocSize int
	key       garray.Array
	keySize   int
	length    []int
	value     []int

	progress     int
	nextCheckPos int

	error_ int

	Searcher
}

func NewDoubleArrayTrie() *DoubleArrayTrie {
	return &DoubleArrayTrie{
		BUF_SIZE:  BUF_SIZE,
		UNIT_SIZE: UNIT_SIZE,
		Searcher: Searcher{
			LongestSearcher: LongestSearcher{
				check: nil,
				base:  nil,
				v:     nil,
			},
		},
		size:      0,
		allocSize: 0,
		error_:    0,
	}
}

// 从TreeMap构造
func NewDoubleArrayTrie2(buildFrom gmap.TreeMap) *DoubleArrayTrie {
	dat := NewDoubleArrayTrie()
	if dat.Build4(buildFrom) != 0 {
		panic("构造失败")
	}
	return dat
}

// 拓展数组
func (dat *DoubleArrayTrie) resize(newSize int) int {
	base2 := make([]int, 0, newSize)
	check2 := make([]int, 0, newSize)
	if dat.allocSize > 0 {
		copy(base2, dat.base)
		copy(check2, dat.check)
	}

	dat.base = base2
	dat.check = check2
	dat.allocSize = newSize
	return dat.allocSize
}

// 获取直接相连的子节点
// parent   父节点
// siblings （子）兄弟节点
// @return 兄弟节点个数
func (dat *DoubleArrayTrie) fetch(parent Node, siblings garray.Array) int {
	if dat.error_ < 0 {
		return 0
	}
	prev := 0
	for i := parent.Left; i < parent.Right; i++ {
		if (dat.length != nil && dat.length[i] < parent.Depth) ||
			(dat.length == nil && len(dat.key.Get(i).(string)) < parent.Depth) {
			continue
		}

		tmp := dat.key.Get(i)
		cur := 0
		if (dat.length != nil && dat.length[i] != parent.Depth) ||
			(dat.length == nil && len(tmp.(string)) != parent.Depth) {
			cur = (String(tmp.(string)).ToCharArray()[parent.Depth]).ToInt() + 1
		}

		if prev > cur {
			dat.error_ = -3
			return 0
		}

		if cur != prev || siblings.Len() == 0 {
			tmp_node := Node{}
			tmp_node.Depth = parent.Depth + 1
			tmp_node.Code = cur
			tmp_node.Left = i
			if siblings.Len() != 0 {
				node := siblings.Get(siblings.Len() - 1).(Node)
				node.Right = i
				siblings.Set(siblings.Len() - 1, node)
			}
			siblings.Append(tmp_node)
		}
		prev = cur
	}

	if siblings.Len() != 0 {
		node := siblings.Get(siblings.Len() - 1).(Node)
		node.Right = parent.Right
		siblings.Set(siblings.Len() - 1, node)
	}

	return siblings.Len()
}

// 插入节点
// siblings 等待插入的兄弟节点
func (dat *DoubleArrayTrie) insert(siblings garray.Array, used bitset.BitSet) int {
	if dat.error_ < 0 {
		return 0
	}

	begin := 0
	pos := gconv.Int(math.Max(gconv.Float64(siblings.Get(0).(Node).Code+1), gconv.Float64(dat.nextCheckPos)) - 1)
	nonzero_num := 0
	first := 0

	if dat.allocSize <= pos {
		dat.resize(pos + 1)
	}

outer:
	// 此循环体的目标是找出满足base[begin + a1...an]  == 0的n个空闲空间,a1...an是siblings中的n个节点
	for {
		pos++
		if dat.allocSize <= pos {
			dat.resize(pos + 1)
		}
		if dat.check[pos] != 0 {
			nonzero_num++
			continue
		} else if first == 0 {
			dat.nextCheckPos = pos
			first = 1
		}

		begin = pos - siblings.Get(0).(Node).Code // 当前位置离第一个兄弟节点的距离
		if dat.allocSize <= (begin + siblings.Get(siblings.Len()-1).(Node).Code) {
			dat.resize(begin + siblings.Get(siblings.Len()-1).(Node).Code + Character_MAX_VALUE)
		}

		if used.Test(gconv.Uint(begin)) {
			continue
		}

		for i := 1; i < siblings.Len(); i++ {
			if dat.check[begin+siblings.Get(i).(Node).Code] != 0 {
				continue outer
			}
		}
		break
	}

	if gconv.Float64(1.0*nonzero_num)/gconv.Float64(pos-dat.nextCheckPos+1) >= 0.95 {
		dat.nextCheckPos = pos // 从位置 next_check_pos 开始到 pos 间，如果已占用的空间在95%以上，下次插入节点时，直接从 pos 位置处开始查找
	}

	used.Set(gconv.Uint(begin))

	if dat.size > begin+siblings.Get(siblings.Len()-1).(Node).Code+1 {
		dat.size = dat.size
	} else {
		dat.size = begin + siblings.Get(siblings.Len()-1).(Node).Code + 1
	}

	for i := 0; i < siblings.Len(); i++ {
		dat.check[begin+siblings.Get(i).(Node).Code] = begin
	}

	for i := 0; i < siblings.Len(); i++ {
		new_siblings := garray.New(true)
		if dat.fetch(siblings.Get(i).(Node), *new_siblings) == 0 { // 一个词的终止且不为其他词的前缀
			if dat.value != nil {
				dat.base[begin+siblings.Get(i).(Node).Code] = -dat.value[siblings.Get(i).(Node).Left] - 1
			} else {
				dat.base[begin+siblings.Get(i).(Node).Code] = -siblings.Get(i).(Node).Left - 1
			}

			if dat.value != nil && -dat.value[siblings.Get(i).(Node).Left]-1 >= 0 {
				dat.error_ = -2
				return 0
			}

			dat.progress++
		} else {
			h := dat.insert(*new_siblings, used)
			dat.base[begin+siblings.Get(i).(Node).Code] = h
		}
	}

	return begin
}

func (dat *DoubleArrayTrie) Clear() {
	dat.check = nil
	dat.base = nil
	dat.allocSize = 0
	dat.size = 0
	dat.error_ = 0
}

func (dat *DoubleArrayTrie) GetUnitSize() int {
	return dat.UNIT_SIZE
}

func (dat *DoubleArrayTrie) GetSize() int {
	return dat.size
}

func (dat *DoubleArrayTrie) GetTotalSize() int {
	return dat.size * dat.UNIT_SIZE
}

func (dat *DoubleArrayTrie) GetNonzeroSize() int {
	result := 0
	for i := 0; i < len(dat.check); i++ {
		if dat.check[i] != 0 {
			result++
		}
	}
	return result
}

func (dat *DoubleArrayTrie) Build(key garray.Array, value garray.Array) int {
	dat.v = value.Slice()
	return dat.Build5(key, nil, nil, key.Len())
}

func (dat *DoubleArrayTrie) Build2(key garray.Array, value []interface{}) int {
	dat.v = value
	return dat.Build5(key, nil, nil, key.Len())
}

/**
 * 构建DAT
 *
 * @param entrySet 注意此entrySet一定要是字典序的！否则会失败
 * @return
 */
func (dat *DoubleArrayTrie) Build3(entrySet []MapEntrySet) int {
	keyList := garray.NewArray(true)
	valueList := garray.New(true)

	for _, entry := range entrySet {
		keyList.Append(entry.GetKey())
		valueList.Append(entry.GetValue())
	}

	return dat.Build(*keyList, *valueList)
}

/**
 * 方便地构造一个双数组trie树
 *
 * @param keyValueMap 升序键值对map
 * @return 构造结果
 */
func (dat *DoubleArrayTrie) Build4(keyValueMap gmap.TreeMap) int {
	entrySetMap := keyValueMap.Map()
	entrySet := utils.MapToMapEntrySet(entrySetMap)
	return dat.Build3(entrySet)
}

/**
 * 唯一的构建方法
 *
 * @param _key     值set，必须字典序
 * @param _length  对应每个key的长度，留空动态获取
 * @param _value   每个key对应的值，留空使用key的下标作为值
 * @param _keySize key的长度，应该设为_key.size
 * @return 是否出错
 */
func (dat *DoubleArrayTrie) Build5(_key garray.Array, _length []int, _value []int, _keySize int) int {
	if &_key == nil || _keySize > _key.Len() {
		return 0
	}

	dat.key = _key
	dat.length = _length
	dat.keySize = _keySize
	dat.value = _value
	dat.progress = 0
	dat.allocSize = 0

	dat.resize(65536 * 32) // 32个双字节

	dat.base[0] = 1
	dat.nextCheckPos = 0

	var root_node Node = Node{}
	root_node.Left = 0
	root_node.Right = dat.keySize
	root_node.Depth = 0

	siblings := garray.New(true)
	dat.fetch(root_node, *siblings)
	dat.insert(*siblings, bitset.BitSet{})
	dat.shrink()

	dat.key = *garray.NewArray(true)
	dat.length = nil

	return dat.error_
}

func (dat *DoubleArrayTrie) Open(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic("get file info error: " + err.Error())
	}
	dat.size = gconv.Int(fileInfo.Size()) / dat.UNIT_SIZE
	dat.check = make([]int, 0, dat.size)
	dat.base = make([]int, 0, dat.size)

	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		panic("read file error: " + err.Error())
	}
	for i := 0; i < dat.size; i++ {
		dat.base[i] = gconv.Int(buf.Bytes()[i])
		dat.check[i] = gconv.Int(buf.Bytes()[i])
	}
}

func (dat *DoubleArrayTrie) Save(fileName string) bool {
	var out bytes.Buffer
	file, _ := os.OpenFile(fileName, os.O_WRONLY, 600)
	defer file.Close()
	out.WriteByte(gconv.Byte(dat.size))
	for i := 0; i < dat.size; i++ {
		out.WriteByte(gconv.Byte(dat.base[i]))
		out.WriteByte(gconv.Byte(dat.check[i]))
	}
	_, err := out.WriteTo(file)
	if err != nil {
		return false
	}
	return true
}

/**
 * 将base和check保存下来
 *
 * @param out
 * @return
 */
func (dat *DoubleArrayTrie) Save2(out *os.File) bool {
	_, err := out.Write(gconv.Bytes(dat.size))
	if err != nil {
		utils.Logger.Warning("write to file err: " + err.Error())
		return false
	}
	for i := 0; i < dat.size; i++ {
		_, err := out.Write(gconv.Bytes(dat.base[i]))
		if err != nil {
			utils.Logger.Warning("write to file err: " + err.Error())
			return false
		}
		_, err = out.Write(gconv.Bytes(dat.check[i]))
		if err != nil {
			utils.Logger.Warning("write to file err: " + err.Error())
			return false
		}
	}

	return true
}

func (dat *DoubleArrayTrie) Save3(out *os.File) {
	_, err := out.Write(gconv.Bytes(dat.base))
	if err != nil {
		utils.Logger.Warning("write to file err: " + err.Error())
		panic(err)
	}
	_, err = out.Write(gconv.Bytes(dat.check))
	if err != nil {
		utils.Logger.Warning("write to file err: " + err.Error())
		panic(err)
	}
}

/**
 * 从磁盘加载，需要额外提供值
 *
 * @param path
 * @param value
 * @return
 */
func (dat *DoubleArrayTrie) Load(path string, value glist.List) bool {
	if !dat.loadBaseAndCheck(path) {
		return false
	}
	dat.v = value.FrontAll()
	return true
}

/**
 * 从磁盘加载，需要额外提供值
 *
 * @param path
 * @param value
 * @return
 */
func (dat *DoubleArrayTrie) Load2(path string, value []interface{}) bool {
	if entry.GConfig.IOAdapter == nil {
		if !dat.loadBaseAndCheckByFileChannel(path) {
			return false
		}
	}else {
		byteArray := ByteArrayStream.CreateByteArrayStream(path)

		if !dat.Load3(*byteArray, value) {
			return false
		}
	}
	dat.v = value
	return true
}


func (dat *DoubleArrayTrie) Load3(byteArray ByteArrayStream.ByteArrayStream, value []interface{}) bool {
	if &byteArray == nil {
		return false
	}

	dat.size = byteArray.NextInt()
	dat.base = make([]int, 0, dat.size + 65535)
	dat.check = make([]int, 0, dat.size + 65535)

	for i := 0; i < dat.size; i++ {
		dat.base[i] = byteArray.NextInt()
		dat.check[i] = byteArray.NextInt()
	}
	dat.v = value
	return true
}

/**
 * 从字节数组加载（发现在MacOS上，此方法比ByteArray更快）
 * @param bytes
 * @param offset
 * @param value
 * @return
 */
func (dat *DoubleArrayTrie) Load4(bytes []byte, offset int, value []interface{}) bool {
	if bytes == nil {
		return false
	}

	dat.size = byteUtility.BytesHighFirstToInt(bytes, offset)
	offset += 4
	dat.base = make([]int, 0, dat.size + 65535)
	dat.check = make([]int, 0, dat.size + 65535)

	for i := 0; i < dat.size; i++ {
		dat.base[i] = byteUtility.BytesHighFirstToInt(bytes, offset)
		offset += 4
		dat.check[i] = byteUtility.BytesHighFirstToInt(bytes, offset)
		offset += 4
	}
	dat.v = value
	return true
}

/**
 * 载入双数组，但是不提供值，此时本trie相当于一个set
 *
 * @param path
 * @return
 */
func (dat *DoubleArrayTrie) Load5(path string) bool {
	return dat.loadBaseAndCheckByFileChannel(path)
}

func (dat *DoubleArrayTrie) loadBaseAndCheck(path string) bool {
	var in *io.DataInputStream
	if entry.GConfig.IOAdapter == nil {
		in = io.NewDataInputStream(path)
	}else{
		in, err := entry.GConfig.IOAdapter.Open(path)
		if err != nil {
			utils.Logger.Warning("open file ", path, " error")
			return false
		}
	}
	dat.size = in.ReadInt()
	dat.base = make([]int, 0, dat.size + 65535)
	dat.check = make([]int, 0, dat.size + 65535)
	for i := 0; i < dat.size; i++ {
		dat.base[i] = in.ReadInt()
		dat.check[i] = in.ReadInt()
	}

	return true
}

func (dat *DoubleArrayTrie) loadBaseAndCheckByFileChannel(path string) bool {

	return true
}

//func (dat *DoubleArrayTrie) SerializeTo(path string) bool {
//
//}
//
//func (dat *DoubleArrayTrie) UnSerialize(path string) *DoubleArrayTrie {
//	return nil
//}

/**
 * 精确匹配
 *
 * @param key 键
 * @return 值
 */
func (dat *DoubleArrayTrie) ExactMatchSearch(key string) int {
	return dat.ExactMatchSearch2(key, 0, 0, 0)
}

func (dat *DoubleArrayTrie) ExactMatchSearch2(key string, pos int, length int, nodePos int) int {
	if length <= 0 {
		length = len(key)
	}
	if nodePos <= 0 {
		nodePos = 0
	}

	result := -1

	b := dat.base[nodePos]
	var p int

	for i := pos; i < length; i++ {
		p = b + String(key).ToCharArray()[i].ToInt() + 1
		if b == dat.check[p] {
			b = dat.base[p]
		} else {
			return result
		}
	}

	p = b
	n := dat.base[p]
	if b == dat.check[p] && n < 0 {
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
func (dat *DoubleArrayTrie) ExactMatchSearch3(keyChars []Character, pos int, length int, nodePos int) int {
	result := -1
	b := dat.base[nodePos]
	var p int

	for i := pos; i < length; i++ {
		p = b + keyChars[i].ToInt() + 1
		if b == dat.check[p] {
			b = dat.base[p]
		} else {
			return result
		}
	}

	p = b
	n := dat.base[p]
	if b == dat.check[p] && n < 0 {
		result = -n - 1
	}
	return result
}

func (dat *DoubleArrayTrie) CommonPrefixSearch(key string) glist.List {
	return dat.CommonPrefixSearch2(key, 0, 0, 0)
}

/**
 * 前缀查询
 *
 * @param key     查询字串
 * @param pos     字串的开始位置
 * @param len     字串长度
 * @param nodePos base中的开始位置
 * @return 一个含有所有下标的list
 */
func (dat *DoubleArrayTrie) CommonPrefixSearch2(key string, pos int, length int, nodePos int) glist.List {
	if length <= 0 {
		length = len(key)
	}

	if nodePos <= 0 {
		nodePos = 0
	}

	result := glist.New(true)

	keyChars := String(key).ToCharArray()

	b := dat.base[nodePos]
	var n, p int

	for i := 0; i < length; i++ {
		p = b + keyChars[i].ToInt() + 1 // 状态转移 p = base[char[i-1]] + char[i] + 1
		if b == dat.check[p] {  // base[char[i-1]] == check[base[char[i-1]] + char[i] + 1]
			b = dat.base[p]
		}else{
			return *result
		}

		p = b
		n = dat.base[p]
		if b == dat.check[p] && n < 0 { // base[p] == check[p] && base[p] < 0 查到一个词
			result.PushBack(-n -1)
		}
	}

	return *result
}

/**
 * 前缀查询，包含值
 *
 * @param key 键
 * @return 键值对列表
 * @deprecated 最好用优化版的
 */
func (dat *DoubleArrayTrie) CommonPrefixSearchWithValue(key string) glist.List {
	length := len(key)
	result := glist.New(true)
	keyChars := String(key).ToCharArray()
	b := dat.base[0]
	var n, p int

	for i := 0; i < length; i++ {
		p = b
		n = dat.base[p]
		if b == dat.check[p] && n < 0 { // base[p] == check[p] && base[p] < 0 查到一个词
			tmp := MapEntrySet{}
			tmp.SetKey(string(keyChars[0:i]))
			tmp.SetValue(dat.v[-n -1])
			result.PushBack(tmp)
		}

		p = b + keyChars[i].ToInt() + 1 // 状态转移 p = base[char[i-1]] + char[i] + 1
		// 下面这句可能产生下标越界，不如改为if (p < size && b == check[p])，或者多分配一些内存
		if b == dat.check[p] {  // base[char[i-1]] == check[base[char[i-1]] + char[i] + 1]
			b = dat.base[p]
		}else{
			return *result
		}
	}

	p = b
	n = dat.base[p]

	if b == dat.check[p] && n < 0 {
		tmp := MapEntrySet{}
		tmp.SetKey(key)
		tmp.SetValue(dat.v[-n -1])
		result.PushBack(tmp)
	}

	return *result
}

/**
 * 优化的前缀查询，可以复用字符数组
 *
 * @param keyChars
 * @param begin
 * @return
 */
func (dat *DoubleArrayTrie) CommonPrefixSearchWithValue2(keyChars []Character, begin int) glist.List {
	length := len(keyChars)
	result := glist.New(true)
	b := dat.base[0]
	var n, p int

	for i:= begin; i<length; i++ {
		p = b
		n = dat.base[p]
		if b == dat.check[p] && n < 0 { // base[p] == check[p] && base[p] < 0 查到一个词
			tmp := MapEntrySet{}
			tmp.SetKey(string(keyChars[begin:i-begin]))
			tmp.SetValue(dat.v[-n -1])
			result.PushBack(tmp)
		}

		p = b + keyChars[i].ToInt() + 1 // 状态转移 p = base[char[i-1]] + char[i] + 1
		// 下面这句可能产生下标越界，不如改为if (p < size && b == check[p])，或者多分配一些内存
		if b == dat.check[p] {  // base[char[i-1]] == check[base[char[i-1]] + char[i] + 1]
			b = dat.base[p]
		}else{
			return *result
		}
	}

	p = b
	n = dat.base[p]

	if b == dat.check[p] && n < 0 {
		tmp := MapEntrySet{}
		tmp.SetKey(string(keyChars[begin:length-begin]))
		tmp.SetValue(dat.v[-n-1])
		result.PushBack(tmp)
	}

	return *result
}

func (dat *DoubleArrayTrie) ToString() string {
	return fmt.Sprintf("DoubleArrayTrie{ size=%d, allocSize=%d, key=%s, keySize=%d, progress=%v, nextCheckPos=%d, error_=%d",
		dat.size, dat.allocSize, dat.key, dat.keySize, dat.progress, dat.nextCheckPos, dat.error_)
}

/**
 * 树叶子节点个数
 *
 * @return
 */
func (dat *DoubleArrayTrie) Size() int {
	return len(dat.v)
}

/**
 * 获取check数组引用，不要修改check
 *
 * @return
 */
func (dat *DoubleArrayTrie) GetCheck() []int {
	return dat.check
}

/**
 * 获取base数组引用，不要修改base
 *
 * @return
 */
func (dat *DoubleArrayTrie) GetBase() []int {
	return dat.base
}

/**
 * 获取index对应的值
 *
 * @param index
 * @return
 */
func (dat *DoubleArrayTrie) GetValueAt(index int) interface{} {
	return dat.v[index]
}

/**
 * 精确查询
 *
 * @param key 键
 * @return 值
 */
func (dat *DoubleArrayTrie) Get(key string) interface{} {
	index := dat.ExactMatchSearch(key)
	if index >= 0 {
		return dat.GetValueAt(index)
	}

	return nil
}

func (dat *DoubleArrayTrie) Get2(key []Character) interface{} {
	index := dat.ExactMatchSearch3(key, 0, len(key), 0)
	if index >= 0 {
		return dat.GetValueAt(index)
	}

	return nil
}

func (dat *DoubleArrayTrie) GetValueArray(a []interface{}) []interface{} {
	size := len(dat.v)
	if len(a) < size {
		a = make([]interface{}, 0, size)
	}
	copy(a, dat.v)
	return a
}

/**
 * 沿着路径转移状态
 *
 * @param path
 * @return
 */
func (dat *DoubleArrayTrie) Transition(path string) int {
	return dat.Transition2(String(path).ToCharArray())
}

/**
 * 沿着节点转移状态
 *
 * @param path
 * @return
 */
func (dat *DoubleArrayTrie) Transition2(path []Character) int {
	b := dat.base[0]
	var p int

	for i := 0; i < len(path); i++ {
		p = b + path[i].ToInt() + 1
		if b == dat.check[p] {
			b = dat.base[p]
		}else{
			return -1
		}
	}

	p = b
	return p
}

/**
 * 沿着路径转移状态
 *
 * @param path 路径
 * @param from 起点（根起点为base[0]=1）
 * @return 转移后的状态（双数组下标）
 */
func (dat *DoubleArrayTrie) Transition3(path string, from int) int {
	b := from
	var p int

	for i:=0; i < len(path); i++ {
		p = b + String(path).ToCharArray()[i].ToInt() + 1
		if b == dat.check[p] {
			b = dat.base[p]
		}else{
			return -1
		}
	}

	p = b
	return p
}

/**
 * 转移状态
 * @param c
 * @param from
 * @return
 */
func (dat *DoubleArrayTrie) Transition4(c Character, from int) int {
	b := from
	var p int

	p = b + c.ToInt() + 1
	if b == dat.check[p] {
		b = dat.base[p]
	}else{
		return -1
	}
	return b
}

/**
 * 检查状态是否对应输出
 *
 * @param state 双数组下标
 * @return 对应的值，null表示不输出
 */
func (dat *DoubleArrayTrie) Output(state int) interface{} {
	if state < 0 {
		return nil
	}
	n := dat.base[state]
	if state == dat.check[state] && n < 0 {
		return dat.v[-n - 1]
	}
	return nil
}

func (dat *DoubleArrayTrie) GetSearcher(text string) Searcher {
	return dat.GetSearcher2(text, 0)
}

func (dat *DoubleArrayTrie) GetSearcher2(text string, offset int) Searcher {
	return *NewSearcher(offset, String(text).ToCharArray())
}

func (dat *DoubleArrayTrie) GetSearcher3(text []Character, offset int) Searcher {
	return *NewSearcher(offset, text)
}

/**
 * 全切分
 *
 * @param text      文本
 * @param processor 处理器
 */
func (dat *DoubleArrayTrie) ParseText(text string, processor ahocorasick.IHit) {
	searcher := dat.GetSearcher2(text, 0)
	for searcher.Next() {
		processor.Hit(searcher.Begin, searcher.Begin + searcher.Length, searcher.Value)
	}
}

func (dat *DoubleArrayTrie) GetLongestSearcher(text string, offset int) LongestSearcher {
	return dat.GetLongestSearcher2(String(text).ToCharArray(), offset)
}

func (dat *DoubleArrayTrie) GetLongestSearcher2(text []Character, offset int) LongestSearcher {
	return *NewLongestSearcher(offset, text)
}

/**
 * 最长匹配
 *
 * @param text      文本
 * @param processor 处理器
 */
func (dat *DoubleArrayTrie) ParseLongestText(text string, processor ahocorasick.IHit) {
	searcher := dat.GetLongestSearcher(text, 0)
	for searcher.Next() {
		processor.Hit(searcher.Begin, searcher.Begin+searcher.Length, searcher.Value)
	}
}

/**
 * 转移状态
 *
 * @param current
 * @param c
 * @return
 */
func (dat *DoubleArrayTrie) Transition5(current int, c Character) int {
	b := dat.base[current]
	var p int

	p = b + c.ToInt() + 1
	if b == dat.check[p] {
		b = dat.base[p]
	} else {
		return -1
	}
	p = b
	return p
}

/**
 * 更新某个键对应的值
 *
 * @param key   键
 * @param value 值
 * @return 是否成功（失败的原因是没有这个键）
 */
func (dat *DoubleArrayTrie) Set(key string, value interface{}) bool {
	index := dat.ExactMatchSearch(key)
	if index >= 0 {
		dat.v[index] = value
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
func (dat *DoubleArrayTrie) Get3(index int) interface{} {
	return dat.v[index]
}

/**
 * 释放空闲的内存
 */
func (dat *DoubleArrayTrie) shrink() {
	nbase := make([]int, 0, dat.size+65535)
	copy(nbase, dat.base)
	dat.base = nbase

	ncheck := make([]int, 0, dat.size+65535)
	copy(ncheck, dat.check)
	dat.check = ncheck
}
