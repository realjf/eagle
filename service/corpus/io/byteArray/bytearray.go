package ByteArray

import (
	. "eagle/service/common"
	"eagle/service/utility/byteUtility"
	"eagle/utils"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
)

// 对字节数组进行封装，提供方便的读取操作
type ByteArray struct {
	Bytes []byte // 当前字节数组，不一定是全部字节，可能只是一个片段
	Offset int // 当前已读取的字节数，或下一个字节的指针
}

func NewByteArray(bytes []byte) *ByteArray {
	return &ByteArray{
		Bytes:bytes,
	}
}

// 从文件读取一个字节数组
func CreateByteArray(path string) *ByteArray {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger.Error("读取文件 ", path, " 失败")
		return nil
	}
	return NewByteArray(bytes)
}

// 获取全部字节
func (ba *ByteArray) GetBytes() []byte {
	return ba.Bytes
}

// 读取一个int
func (ba *ByteArray) NextInt() int {
	result := byteUtility.BytesHighFirstToInt(ba.Bytes, ba.Offset)
	ba.Offset += 4
	return result
}

func (ba *ByteArray) NextFloat64() float64 {
	result := byteUtility.BytesHighFirstToFloat64(ba.Bytes, ba.Offset)
	result += 8
	return result
}

// 读取一个char，对应于writeChar
func (ba *ByteArray) NextChar() Char {
	result := byteUtility.BytesHighFirstToChar(ba.Bytes, ba.Offset)
	ba.Offset += 2
	return result
}

// 读取一个字节
func (ba *ByteArray) NextByte() byte {
	result := ba.Bytes[ba.Offset]
	ba.Offset++
	return result
}

//
func (ba *ByteArray) NextBoolean() bool {
	return ba.NextByte() == 1
}


func (ba *ByteArray) HasMore() bool {
	return ba.Offset < len(ba.Bytes)
}

func (ba *ByteArray) NextString() string {
	var buf []Char = make([]Char, ba.NextInt())
	for i := 0; i < len(buf); i++ {
		buf[i] = ba.NextChar()
	}
	s := NewStringFromChar(buf)
	return s.ToString()
}

func (ba *ByteArray) NextFloat() float32 {
	result := byteUtility.BytesHighFirstToFloat(ba.Bytes, ba.Offset)
	ba.Offset += 4
	return result
}

// 读取一个无符号短整型
func (ba *ByteArray) NextUnsignedShort() int {
	var a byte = ba.NextByte()
	var b byte = ba.NextByte()
	return gconv.Int(((a & 0xff) << 8) | (b & 0xff))
}

// 读取一个UTF字符串
func (ba *ByteArray) NextUTF() string {
	var utflen int = ba.NextUnsignedShort()
	var bytearr []byte = make([]byte, utflen)
	var chararr []Char = make([]Char, utflen)

	var c, char2, char3 int
	count := 0
	chararr_count := 0

	for i := 0; i < utflen; i++ {
		bytearr[i] = ba.NextByte()
	}

	for count < utflen {
		c = int(bytearr[count] & 0xff)
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = Char(c)
		chararr_count++
	}

	for count < utflen {
		c = int(bytearr[count] & 0xff)
		switch c >> 4 {
		case 0: fallthrough
		case 1: fallthrough
		case 2: fallthrough
		case 3: fallthrough
		case 4: fallthrough
		case 5: fallthrough
		case 6: fallthrough
		case 7:
			count++
			chararr[chararr_count] = Char(c)
			chararr_count++
		case 12: fallthrough
		case 13:
			count += 2
			if count > utflen {
				utils.Logger.Info("malformed input: partial character at end")
			}
			char2 = int(bytearr[count -1])
			if char2 & 0xc0 != 0x80 {
				utils.Logger.Info("malformed input around byte ", count)
			}
			chararr[chararr_count] = Char(((c & 0x1F) << 6) | (char2 & 0x3F))
			chararr_count++
		case 14:
			count += 3
			if count > utflen {
				utils.Logger.Info("malformed input: partial character at end")
			}
			char2 = int(bytearr[count - 2])
			char3 = int(bytearr[count - 1])
			if (char2 & 0xC0) != 0x80 || (char3 & 0xC0) != 0x80 {
				utils.Logger.Info("malformed input around byte ", count-1)
			}
			chararr[chararr_count] = Char( ((c & 0x0F) << 12) | ((char2 & 0x3F) << 6) | ((char3 & 0x3F) << 0) )
			chararr_count++
		default:
			utils.Logger.Info("malformed input around byte ", count)
		}
	}
	s := NewStringFromChar(chararr[0:chararr_count])
	return s.ToString()
}


func (ba *ByteArray) GetOffset() int {
	return ba.Offset
}

func (ba *ByteArray) GetLength() int {
	return len(ba.Bytes)
}

func (ba *ByteArray) Close() {
	ba.Bytes = nil
}





