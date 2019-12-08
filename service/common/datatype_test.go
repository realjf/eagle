package common

import (
	"testing"
)

func TestCharacter_ToInt(t *testing.T) {
	var s = 'c'
	t.Log(Char(s).ToInt())
	var b = "c"
	t.Log(String(b).ToInt())
	t.FailNow()
}

func TestCharacter_ToString(t *testing.T) {
	var s = 'c'
	t.Fatal(Char(s).ToString())
}

func TestString_ToCharArray(t *testing.T) {
	var s string = "hwojo给是个个奇偶陪我今儿去骗人１４２３４１ ee"
	t.Log(String(s).ToStringArray())
	t.Fatal(String(s).ToCharArray())
}

func TestString_Length(t *testing.T) {
	var s string = "hwojoee"
	t.Fatal(String(s).Length())
}

func TestCharacter_ToByte(t *testing.T) {
	var s string = "geqwer我"
	t.Log(String(s).CharAt(6).ToByte())
	t.Fatal(String(s).CharAt(2).ToByte())
}
