package byteUtility

import (
	"math/big"
	"testing"
)

func TestBytesToChar(t *testing.T) {
	var bytes []byte = []byte(string("hell134143o"))
	var l int64 = int64(int64(bytes[0]) << 56)
	l1 := big.NewInt(0)
	l1.And(big.NewInt(l), big.NewInt(0xFF0000000000000))
	t.Log(uint64(l) & uint64(0xFF0000000000000))
	t.Log(BytesToInt64(bytes))
	t.Logf("%s", Int64ToBytes(BytesToInt64(bytes)))
	t.Fatal(l1)
}

func TestBytesHighFirstToChar(t *testing.T) {
	var bytes []byte = []byte(string("hell134143o"))
	t.Log(BytesHighFirstToChar(bytes, 0))
	t.FailNow()
}

func TestBytesHighFirstToInt64(t *testing.T) {
	var bytes []byte = []byte(string("hell134143o"))
	t.Log(BytesHighFirstToInt64(bytes))
	t.Logf("%s", Int64ToBytes(BytesHighFirstToInt64(bytes)))
	t.FailNow()
}

func TestBytesHighFirstToInt(t *testing.T) {
	var bytes []byte = []byte(string("hell134143o"))
	t.Log(BytesHighFirstToInt(bytes, 0))
	t.Logf("%s", IntToBytes(BytesHighFirstToInt(bytes, 0)))
	t.FailNow()
}

