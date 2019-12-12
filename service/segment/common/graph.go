package common

import (
	"fmt"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/util/gconv"
)

type Graph struct {
	Vertexes []Vertex // 顶点
	EdgesTo []*glist.List // 边，到达下标i
}

func NewGraph(vertexes []Vertex) *Graph {
	g := &Graph{
		Vertexes:vertexes,
	}
	size := len(vertexes)
	g.EdgesTo = make([]*glist.List, size)
	for i := 0; i < size; i++ {
		g.EdgesTo[i] = glist.New(true)
	}
	return g
}

/**
 * 连接两个节点
 * @param from 起点
 * @param to 终点
 * @param weight 花费
 */
func (g *Graph) Connect(from int, to int, weight float64) {
	g.EdgesTo[to].PushBack(*NewEdgeFrom(from, weight, g.Vertexes[from].Word + "@" + g.Vertexes[to].Word))
}

/**
 * 获取到达顶点to的边列表
 * @param to 到达顶点to
 * @return 到达顶点to的边列表
 */
func (g *Graph) GetEdgeListTo(to int) glist.List {
	return *g.EdgesTo[to]
}

func (g *Graph) ToString() string {
	return fmt.Sprintf("Graph{vertexes=%s, edgesTo=%s}", gconv.SliceStr(g.Vertexes), gconv.SliceStr(g.EdgesTo))
}

func (g *Graph) PrintByTo() string {
	sb := "========按终点打印========"
	for to := 0; to < len(g.EdgesTo); to++ {
		edgeFromList := g.EdgesTo[to]
		for _, edgeFrom := range edgeFromList.FrontAll() {
			eFrom := edgeFrom.(EdgeFrom)
			sb += fmt.Sprintf("to:%3d, from:%3d, weight:%05.2f, word:%s\n", to, eFrom.From, eFrom.Weight, eFrom.Name)
		}
	}
	return sb
}

/**
 * 根据节点下标数组解释出对应的路径
 * @param path
 * @return
 */
func (g *Graph) ParsePath(path []int) glist.List {
	vertexList := glist.New(true)
	for _,i := range path {
		vertexList.PushBack(g.Vertexes[i])
	}

	return *vertexList
}

func (g *Graph) GetVertexes() []Vertex {
	return g.Vertexes
}

func (g *Graph) GetEdgesTo() []*glist.List {
	return g.EdgesTo
}
