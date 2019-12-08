package ahocorasick

import "github.com/gogf/gf/container/gmap"

type Builder struct {
	rootState State // 根节点，仅仅用于构建过程
	used []bool // 是否占用，仅仅用于构建
	allocSize int // 已分配在内存中的大小
	progress int // 一个控制增长速度的变量
	nextCheckPos int // 下一个插入的位置将从此开始搜索
	keySize int // 键值对的大小
}

func (b *Builder) Build(ma gmap.TreeMap) {

}




