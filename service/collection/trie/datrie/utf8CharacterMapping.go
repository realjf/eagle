package datrie

import (
	"eagle/service/common"
)

//  UTF-8编码到int的映射
type Utf8CharacterMapping struct {
	serialVersionUID int64
	N int
	EMPTYLIST []int
	UTF_8 string
}

func (u *Utf8CharacterMapping) GetInitSize() int {
	return u.N
}

func (u *Utf8CharacterMapping) GetCharsetSize() int {
	return u.N
}

func (u *Utf8CharacterMapping) ZeroId() int {
	return 0
}

func (u *Utf8CharacterMapping) ToIdList(key string) []int {
	bytes
	res := make([]int, )
	for i := 0; i < len(res); i++ {
		res[i] = bytes[i] & 0xFF
	}
	if len(res) == 1 && res[0] == 0 {
		return u.EMPTYLIST
	}

	return res
}

func (u *Utf8CharacterMapping) ToIdList2(codePoint int) []int {
	var count int
	if codePoint < 0x80 {
		count = 1
	}else if codePoint < 0x800 {
		count = 2
	}else if codePoint < 0x10000 {
		count = 3
	}else if codePoint < 0x200000 {
		count = 4
	}else if codePoint < 0x4000000 {
		count = 5
	}else if codePoint < 0x7fffffff {
		count = 6
	}else {
		return u.EMPTYLIST
	}

	r := []int{}
	switch count {
	case 6:
		r[5] = common.Char(0x80 | (codePoint & 0x3f)).ToInt()
		codePoint = codePoint >> 6
		codePoint |= 0x4000000
	case 5:
		r[4] = common.Char(0x80 | (codePoint & 0x3f)).ToInt()
		codePoint = codePoint >> 6
		codePoint |= 0x200000
	case 4:
		r[3] = common.Char(0x80 | (codePoint & 0x3f)).ToInt()
		codePoint = codePoint >> 6
		codePoint |= 0x10000
	case 3:
		r[2] = common.Char(0x80 | (codePoint & 0x3f)).ToInt()
		codePoint = codePoint >> 6
		codePoint |= 0x800
	case 2:
		r[1] = common.Char(0x80 | (codePoint & 0x3f)).ToInt()
		codePoint = codePoint >> 6
		codePoint |= 0xc0
	case 1:
		r[0] = common.Char(codePoint).ToInt()
	}
	return r
}

func (u *Utf8CharacterMapping) ToString(ids []int) string {
	var bytes []byte = make([]byte, len(ids))
	for i := 0; i < len(ids); i++ {
		bytes[i] = byte(ids[i])
	}
	return string(bytes)
}

