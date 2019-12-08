package ns

// 地名角色标签
type NS int

const (
	A NS = iota // 地名的上文 我【来到】中关园
	B // 地名的下文刘家村/和/下岸村/相邻
	C // 中国地名的第一个字
	D // 中国地名的第二个字
	E // 中国地名的第三个字
	G // 其他整个的地名
	H // 中国地名的后缀海/淀区
	X // 连接词刘家村/和/下岸村/相邻
	Z // 其它非地名成分
	S // 句子的开头
)

