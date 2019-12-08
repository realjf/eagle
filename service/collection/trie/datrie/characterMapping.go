package datrie

type CharacterMapping interface {
	GetInitSize() int
	GetCharsetSize() int
	ZeroId() int
	ToIdList(key string) []int
	ToIdList2(codePoint int) []int
	ToString(ids []int) string
}


