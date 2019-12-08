package TextUtility

import (
	"testing"
)

func TestIsAllChinese(t *testing.T) {
	str := "你好，我是中国人"
	str1 := "我是jioge"
	t.Fatal(IsAllChinese(str), IsAllNonChinese([]byte(str1)))
}

func TestCharTypes(t *testing.T) {
	str := "我"
	t.Fatal(CharTypes(str))
}

func TestIsAllNum(t *testing.T) {
	str := "1234134"
	t.Log(IsAllNum(str))
	str1 := "百"
	t.Log(IsAllNum(str1))
	str2 := "０１１３４１３４"
	t.Log(IsAllNum(str2))
	str3 := "＋1234"
	t.Log(IsAllNum(str3))
	str4 := "f32323"
	t.Fatal(IsAllNum(str4))
}

func TestIsAllLetter(t *testing.T) {
	str := "12314"
	t.Log(IsAllLetter(str))
	str1 := "jigoeej"
	t.Fatal(IsAllLetter(str1))
}

func TestIsAllLetterOrNum(t *testing.T) {
	str := "geojie13424"
	t.Log(IsAllLetterOrNum(str))
	str2 := "23414"
	t.Log(IsAllLetterOrNum(str2))
	str3 := "几个宁儿"
	t.Log(IsAllLetterOrNum(str3))
	str1 := "jignegje"
	t.Fatal(IsAllLetterOrNum(str1))
}


