package interval

type Intervalable interface {
	// 起点
	GetStart() int
	// 终点
	GetEnd() int
	// 长度
	Size() int
}
