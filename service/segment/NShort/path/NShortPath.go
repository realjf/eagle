package path

import (
	. "eagle/service/segment/common"
	"eagle/service/utility/predefine"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/util/gconv"
	"math"
)

type NShortPath struct {
	graph Graph
	n int
	vertexCount int
	fromArray [][]CQueue
	weightArray [][]float64
}

/**
 * 构造一个N最短路径计算器
 * @param graph 要计算的图
 * @param N 要计算前几条最短路径，当然结果不一定就是N条
 */
func NewNShortPath(graph Graph, n int) *NShortPath {
	ns := &NShortPath{}
	ns.calculate(graph, n)
	return ns
}

/**
 * 初始化，主要分配内存
 * @param inGraph 输入图
 * @param nValueKind 希望的N值
 */
func (n *NShortPath) initNShortPath(inGraph Graph, nValueKind int) {
	n.graph = inGraph
	n.n = nValueKind

	n.vertexCount = len(inGraph.Vertexes)

	n.fromArray = make([][]CQueue, 0, n.vertexCount-1)
	n.weightArray = make([][]float64, 0, n.vertexCount-1)

	for i := 0; i < n.vertexCount - 1; i++ {
		n.fromArray[i] = make([]CQueue, 0, nValueKind)
		n.weightArray[i] = make([]float64, 0, nValueKind)

		for j := 0; j < nValueKind; j++ {
			n.fromArray[i][j] = *NewCQueue()
		}
	}
}

/**
 * 计算出所有结点上可能的路径，为路径数据提供数据准备
 * @param inGraph 输入图
 * @param nValueKind 前N个结果
 */
func (n *NShortPath) calculate(inGraph Graph, nValueKind int) {
	n.initNShortPath(inGraph, nValueKind)

	var tmpElement QueueElement
	queWork := NewCQueue()
	var eWeight float64

	for nCurNode := 1; nCurNode < n.vertexCount; nCurNode++ {
		// 将所有到当前结点（nCurNode)可能到达的边根据eWeight排序并压入队列
		n.enQueueCurNodeEdges(*queWork, nCurNode)

		// 初始化当前结点所有边的eWeight值
		for i := 0; i < n.n; i++ {
			n.weightArray[nCurNode-1][i] = math.MaxFloat64
		}

		// 将queWork中的内容装入fromArray
		tmpElement = queWork.DeQueue()
		if &tmpElement != nil {
			for i :=0; i < n.n; i++ {
				eWeight = tmpElement.Weight
				n.weightArray[nCurNode-1][i] = eWeight

				for tmpElement.Weight == eWeight {
					n.fromArray[nCurNode-1][i].EnQueue(*NewQueueElement(tmpElement.From, tmpElement.Index, 0))
					tmpElement = queWork.DeQueue()
					if &tmpElement == nil {
						i = n.n
						break
					}
				}
			}
		}
	}
}


/**
 * 将所有到当前结点（nCurNode）可能的边根据eWeight排序并压入队列
 * @param queWork
 * @param nCurNode
 */
func (n *NShortPath) enQueueCurNodeEdges(queWork CQueue, nCurNode int) {
	var nPreNode int
	var eWeight float64
	var pEdgeToList glist.List

	queWork.Clear()
	pEdgeToList = n.graph.GetEdgeListTo(nCurNode)

	for _, e := range pEdgeToList.FrontAll() {
		ef := e.(EdgeFrom)
		nPreNode = ef.From
		eWeight = ef.Weight

		for i := 0; i < n.n; i++ {
			// 第一个结点，没有PreNode，直接加入队列
			if nPreNode == 0 {
				queWork.EnQueue(*NewQueueElement(nPreNode, i, eWeight))
				break
			}

			// 如果PreNode的Weight == INFINITE_VALUE，则没有必要继续下去了
			if n.weightArray[nPreNode - 1][i] == math.MaxFloat64 {
				break
			}
			queWork.EnQueue(*NewQueueElement(nPreNode, i, eWeight + n.weightArray[nPreNode-1][i]))
		}
	}
}


/**
 * 获取前index+1短的路径
 * @param index index ＝ 0 : 最短的路径； index = 1 ： 次短的路径, 依此类推。index <= this.N
 * @return
 */
func (n *NShortPath) GetPaths(index int) glist.List {
	stack := garray.New(true)
	curNode := n.vertexCount - 1
	curIndex := index
	var element QueueElement
	var node PathNode
	var aPath []int
	result := glist.New(true)

	element = n.fromArray[curNode-1][curIndex].GetFirst()
	for &element != nil {
		// ---------- 通过压栈得到路径 -----------
		stack.PushLeft(*NewPathNode(curNode, curIndex))
		stack.PushLeft(*NewPathNode(element.From, element.Index))
		curNode = element.From

		for curNode != 0 {
			element = n.fromArray[element.From -1][element.Index].GetFirst()
			stack.PushLeft(*NewPathNode(element.From, element.Index))
			curNode = element.From
		}

		// -------------- 输出路径 --------------
		nArray := make([]PathNode, 0, stack.Len())
		for i:=0; i<stack.Len(); i++ {
			nArray[i] = stack.Get(stack.Len()- i -1).(PathNode)
		}
		aPath = make([]int, 0, len(nArray))

		for i :=0; i < len(aPath); i++ {
			aPath[i] = nArray[i].From
		}
		result.PushBack(aPath)

		// -------------- 出栈以检查是否还有其它路径 --------------
		node = stack.PopLeft().(PathNode)
		curNode = node.From
		curIndex = node.Index
		for curNode < 1 || (stack.Len() != 0 && !n.fromArray[curNode-1][curIndex].CanGetNext()) {
			node = stack.PopLeft().(PathNode)
			curNode = node.From
			curIndex = node.Index
		}
		element = n.fromArray[curNode-1][curIndex].GetNext()
	}

	return *result
}

/**
 * 从短到长获取至多 n 条路径
 * @param n
 * @return
 */
func (ns *NShortPath) GetNPaths(n int) glist.List {
	result := glist.New(true)
	n = gconv.Int(math.Min(float64(predefine.MAX_SEGMENT_NUM), float64(n)))
	for i := 0; i < ns.n && result.Size() < n; i++ {
		pathList := ns.GetPaths(i)
		for _, path := range pathList.FrontAll() {
			p := path.([]int)
			if result.Size() == n {
				break
			}
			result.PushBack(p)
		}
	}
	return *result
}
