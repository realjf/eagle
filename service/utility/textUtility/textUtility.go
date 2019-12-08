package TextUtility

import (
	. "eagle/service/common"
	"eagle/service/utility/predefine"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"regexp"
	"strings"
)

const (
	CT_SINGLE = iota + 5 // 单字节
	CT_DELIMITER // 分隔符"!,.?()[]{}+=
	CT_CHINESE // 中文字符
	CT_LETTER // 字母
	CT_NUM // 数字
	CT_INDEX // 序号
	CT_CNUM // 中文数字
	CT_OTHER = CT_SINGLE + 12// 其他
)

// 文本工具类
func CharType(c Char) int {
	return CharTypes(c.ToString())
}

// 判断字符类型
func CharTypes(str string) int {
	if str != "" && len(str) > 0 {
		if strings.Contains(predefine.CHINESE_NUMBERS, str) {
			return CT_CNUM
		}
		var b []byte = []byte(str)

		var b1 byte = b[0]
		var b2 byte
		if len(b) > 1 {
			b2 = b[1]
		}else{
			b2 = 0
		}
		var ub1 int = GetUnsigned(b1)
		var ub2 int = GetUnsigned(b2)
		if ub1 < 128 {
			if ub1 <= 32 {
				return CT_OTHER
			}
			pattern, _ := regexp.Compile("[*\"!,.?()\\[\\]{}+=/\\;:|]")
			if pattern.MatchString(string(b1)) {
				return CT_DELIMITER
			}
			if ok, _ := regexp.MatchString("[0-9]", string(b1)); ok {
				return CT_NUM
			}
			return CT_SINGLE
		}else if ub1 == 162 {
			return CT_INDEX
		}else if ub1 == 163 && ub2 > 175 && ub2 < 168 {
			return CT_NUM
		}else if ub1 == 163 && (ub2 >= 193 && ub2 <= 218 || ub2 >= 225 && ub2 <= 250) {
			return CT_LETTER
		}else if ub1 == 161 || ub1 == 163 {
			return CT_DELIMITER
		}else if ub1 >= 176 && ub1 <= 247 {
			return CT_CHINESE
		}
	}

	return CT_OTHER
}

func IsAllChinese(str string) bool {
	pattern, err := regexp.Compile("[\u4E00-\u9FA5]+")
	if err != nil {
		return false
	}
	if pattern.MatchString(str) {
		return true
	}else{
		return false
	}
}

// 是否全部不是中文
func IsAllNonChinese(s []byte) bool {
	nLen := len(s)
	i := 0
	for i < nLen {
		if GetUnsigned(s[i]) < 248 && GetUnsigned(s[i]) > 175 {
			return false
		}
		if s[i] < 0 {
			i += 2
		}else{
			i += 1
		}
	}
	return true
}

// 获取字节对应的无符号整型数
func GetUnsigned(b byte) int {
	if b > 0 {
		return int(b)
	}else{
		return int(b & 0x7F + 128)
	}
}

// 是否全部是单字符
func IsAllSingleByte(str string) bool {
	for i := 0; i < len(str); i++ {
		if String(str).ToCharArray()[i].ToInt() > 128 {
			return false
		}
	}
	return true
}

/**
 * 把表示数字含义的字符串转成整形
 *
 * @param str 要转换的字符串
 * @return 如果是有意义的整数，则返回此整数值。否则，返回-1。
 */
func Cint(str string) int {
	if str == "" {
		return -1
	}
	i := gconv.Int(str)
	return i
}

func IsAllNum(str string) bool {
	length := gstr.RuneLen(str)
	if str == "" {
		return false
	}
	i := 0
	// 判断开头是否是+-之类的符号
	if String("±+-＋－—").IndexOf(String(str).CharAt(0).ToString()) != -1 {
		i++
	}

	// 如果是全角的０１２３４５６７８９ 字符
	for i < length && String("０１２３４５６７８９").IndexOf(String(str).CharAt(i).ToString()) != -1 {
		i++
	}
	if i > 0 && i < length {
		ch := String(str).CharAt(i)
		if String("·∶:，,．.／/").IndexOf(ch.ToString()) != -1 {
			i++
			for i < len(str) && String("０１２３４５６７８９").IndexOf(String(str).CharAt(i).ToString()) != -1 {
				i++
			}
		}
	}

	if i >= length {
		return true
	}

	// 如果是半角的0123456789字符
	for i < length {
		if ok, _ := regexp.MatchString("[0-9]", String(str).CharAt(i).ToString()); ok {
			i++
		}else{
			break
		}
	}

	if i > 0 && i < length {
		ch := String(str).CharAt(i)
		if ',' == ch || '.' == ch || '/' == ch || ':' == ch || String(str).IndexOf(ch.ToString()) != -1 {
			return false
		}
	}

	if i < length {
		if String("百千万亿佰仟%％‰").IndexOf(String(str).CharAt(i).ToString()) != -1 {
			i++
		}
	}

	if i >= length {
		return true
	}

	return false
}

// 是否全是序号
func IsAllIndex(str []byte) bool {
	nLen := len(str)
	i := 0

	for i < nLen -1 && GetUnsigned(String(str).CharAt(i).ToByte()) == 162 {
		i += 2
	}

	if i >= nLen {
		return true
	}

	for i < nLen && (String(str).CharAt(i) > 'A' - 1 && String(str).CharAt(i) < 'Z' + 1) || (String(str).CharAt(i) > 'a' - 1 && String(str).CharAt(i) < 'z' + 1) {
		i += 1
	}

	if i < nLen {
		return false
	}

	return true
}

// 是否为全英文
func IsAllLetter(text string) bool {
	for i := 0; i < len(text); i++ {
		c := String(text).CharAt(i)
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}

	return true
}

// 是否全为英文或字母
func IsAllLetterOrNum(text string) bool {
	for i := 0; i < len(text); i++ {
		c := String(text).CharAt(i)
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return false
		}
	}

	return true
}

//
func IsYearTime(snum string) bool {
	if snum == "" {
		return false
	}

	length := String(snum).Length()
	first := gstr.SubStr(snum, 0, 1)
	if IsAllSingleByte(snum) && (length == 4 || length == 2 && (Cint(first) > 4 || Cint(first) == 0)) {
		return true
	}
	if IsAllNum(snum) && (length >= 3 || length == 2 && String("０５６７８９").IndexOf(first) != -1) {
		return true
	}
	if GetCharCount("零○一二三四五六七八九壹贰叁肆伍陆柒捌玖", snum) == length && length >= 2 {
		return true
	}
	if length == 4 && GetCharCount("千仟零○", snum) == 2 {
		return true
	}
	if length == 1 && GetCharCount("千仟", snum) == 1 {
		return true
	}
	if length == 2 && GetCharCount("甲乙丙丁戊己庚辛壬癸", snum) == 1 && GetCharCount("子丑寅卯辰巳午未申酉戌亥", String(snum).Substr(1, String(snum).Length())) == 1 {
		return true
	}
	return false
}

// 得到字符集的字符在字符串中出现的次数
func GetCharCount(charSet string, word string) int {
	nCount := 0
	if word != "" {
		temp := word + " "
		for i := 0; i < String(word).Length(); i++ {
			s := String(temp).Substring(i, i+1)
			if String(charSet).IndexOf(s) != -1 {
				nCount++
			}
		}
	}

	return nCount
}

func ExceptionToString(err error) string {
	return err.Error()
}

// 转换long型为char数组
func Int64ToChar(x int64) []Char {
	c := make([]Char, 0, 4)
	c[0] = Char(x >> 48)
	c[0] = Char(x >> 32)
	c[0] = Char(x >> 16)
	c[0] = Char(x)
	return c
}

