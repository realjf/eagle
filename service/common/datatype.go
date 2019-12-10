package common

import (
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"unicode"
)

var (
	Character_MAX_VALUE = gconv.Int(rune('\uFFFF'))
)

type Char rune

type String string

type Rune rune

type V interface{}

func NewString() String {
	return String("")
}

// 转化为utf-8字符数组
func (s String) ToCharArray() []Char {
	str := strings.Split(string(s), "")
	res := []Char{}
	for _, ch := range str {
		res = append(res, Char(gconv.Runes(ch)[0]))
	}
	return res
}

func (s String) ToStringArray() []string {
	return strings.Split(string(s), "")
}

func (s String) Substring(start int, end int) string {
	length := end - start
	return gstr.SubStr(string(s), start, length)
}

func (s String) LastIndexOf(needle string, startOffset ...int) int {
	return gstr.PosR(string(s), needle, startOffset...)
}

func (s String) StartsWith(prefix string) bool {
	prefixs := String(prefix).ToCharArray()
	value := s.ToCharArray()[0]
	for _, pre := range prefixs {
		if pre == value {
			return true
		}
	}
	return false
}

func (s String) Substr(start int, length int) string {
	return gstr.SubStr(string(s), start, length)
}

func (s String) Length() int {
	return len(strings.Split(string(s), ""))
}

func (s String) CharAt(i int) Char {
	return s.ToCharArray()[i]
}

func (s String) ToInt() int {
	if len(s) == 1 {
		return gconv.Int([]rune(s)[0])
	}else{
		return gconv.Int(s)
	}
}

func (s String) ToBytes() []byte {
	return gconv.Bytes(string(s))
}

func (s *String) Append(v ...interface{}) *String {
	sb := ""
	for _, val := range v {
		sb += gconv.String(val)
	}
	*s = String(string(*s)+sb)
	return s
}

func (s *String) ToString() string {
	return string(*s)
}

func (s Rune) ToString() string {
	return gconv.String(string(s))
}

func (s String) IndexOf(needle string) int {
	return gstr.Pos(string(s), needle)
}

func (s String) IndexOfI(needle string) int {
	return gstr.PosI(string(s), needle)
}

func (c Char) ToString() string {
	return gconv.String(string(c))
}

func (c Char) ToInt() int {
	return gconv.Int(c)
}

func (c Char) ToByte() byte {
	return gconv.Byte(rune(c))
}

func (c Char) ToRune() rune {
	return rune(c)
}

func (c Char) IsWhitespace() bool {
	return unicode.IsSpace(c.ToRune())
}

type MapEntrySet map[string]interface{}

func (s *MapEntrySet) GetKey() interface{} {
	return (*s)["key"]
}

func (s *MapEntrySet) GetValue() interface{} {
	return (*s)["value"]
}

func (s *MapEntrySet) SetKey(key interface{}) {
	(*s)["key"] = key
}

func (s *MapEntrySet) SetValue(value interface{}) {
	(*s)["value"] = value
}

type EntrySet map[interface{}]interface{}

func (s *EntrySet) GetKey() interface{} {
	return (*s)["key"]
}

func (s *EntrySet) GetValue() interface{} {
	return (*s)["value"]
}

func (s *EntrySet) SetKey(key interface{}) {
	(*s)["key"] = key
}

func (s *EntrySet) SetValue(value interface{}) {
	(*s)["value"] = value
}
