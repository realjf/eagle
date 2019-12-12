/**
 * 对数字和字节进行转换。<br>
 * 基础知识：<br>
 * 假设数据存储是以大端模式存储的：<br>
 * byte: 字节类型 占8位二进制 00000000<br>
 * char: 字符类型 占2个字节 16位二进制 byte[0] byte[1]<br>
 * int : 整数类型 占4个字节 32位二进制 byte[0] byte[1] byte[2] byte[3]<br>
 * long: 长整数类型 占8个字节 64位二进制 byte[0] byte[1] byte[2] byte[3] byte[4] byte[5]
 * byte[6] byte[7]<br>
 * float: 浮点数(小数) 占4个字节 32位二进制 byte[0] byte[1] byte[2] byte[3]<br>
 * double: 双精度浮点数(小数) 占8个字节 64位二进制 byte[0] byte[1] byte[2] byte[3] byte[4]
 * byte[5] byte[6] byte[7]<br>
 */

package byteUtility

import (
	. "eagle/service/common"
	"github.com/gogf/gf/util/gconv"
)

/**
 * 将一个2位字节数组转换为char字符。<br>
 * 注意，函数中不会对字节数组长度进行判断，请自行保证传入参数的正确性。
 *
 * @param b 字节数组
 * @return char字符
 */
func BytesToChar(b []byte) Char {
	var c Char = Char(gconv.Int64(b[0] << 8) & 0xFF00)
	c |= Char(b[1] & 0xFF)
	return c
}


/**
 * 字节数组和整型的转换，高位在前，适用于读取writeInt的数据
 *
 * @param bytes 字节数组
 * @return 整型
 */
func BytesHighFirstToInt(bb []byte, start int) int {
	var bytes []byte = make([]byte, 4)
	copy(bytes, bb)
	var num int = gconv.Int(uint(bytes[start+3]) & 0xFF)
	num |= gconv.Int(uint(bytes[start+2]) << 8 ) & 0xFF00
	num |= gconv.Int(uint(bytes[start+1]) << 16 ) & 0xFF0000
	num |= gconv.Int(uint(bytes[start]) << 24 ) & 0xFF000000
	return num
}

/**
 * 字节数组转char，高位在前，适用于读取writeChar的数据
 *
 * @param bytes
 * @param start
 * @return
 */
func BytesHighFirstToChar(bytes []byte, start int) Char {
	var c Char = Char((bytes[start] & 0xFF) << 8 | bytes[start+1] & 0xFF)
	return c
}

// 读取float64，高位在前
func BytesHighFirstToFloat64(bb []byte, start int) float64 {
	var bytes []byte = make([]byte, 8)
	copy(bytes, bb)
	var l int64 = int64(uint64(bytes[start]) << 56)
	l = int64(uint64(l) & uint64(0xFF00000000000000))
	// 如果不强制转换为long，那么默认会当作int，导致最高32位丢失
	l |= int64(uint64(bytes[start + 1]) << 48) & 0xFF000000000000
	l |= int64(uint64(bytes[start + 2]) << 40) & 0xFF0000000000
	l |= int64(uint64(bytes[start + 3]) << 32) & 0xFF00000000
	l |= int64(uint64(bytes[start + 4]) << 24) & 0xFF000000
	l |= int64(uint64(bytes[start + 5]) << 16) & 0xFF0000
	l |= int64(uint64(bytes[start + 6]) << 8) & 0xFF00
	l |= int64(uint64(bytes[start + 7])) & 0xFF

	return gconv.Float64(l)
}


func BytesHighFirstToFloat(bytes []byte, start int) float32 {
	var l int = BytesHighFirstToInt(bytes, start)
	return gconv.Float32(l)
}

/**
 * 将一个4位字节数组转换为4整数。<br>
 * 注意，函数中不会对字节数组长度进行判断，请自行保证传入参数的正确性。
 *
 * @param b 字节数组
 * @return 整数
 */
func BytesToInt(bb []int) int {
	var b []int = make([]int, 4)
	copy(b, bb)
	var i int = (b[0] << 24) & 0xFF000000
	i |= (b[1] << 16) & 0xFF0000
	i |= (b[2]  << 8) & 0xFF00
	i |= b[3] & 0xFF
	return i
}

/**
 * 将一个8位字节数组转换为长整数。<br>
 * 注意，函数中不会对字节数组长度进行判断，请自行保证传入参数的正确性。
 *
 * @param b 字节数组
 * @return 长整数
 */
func BytesToInt64(bb []byte) int64 {
	var b []byte = make([]byte, 8)
	copy(b, bb)
	var l int64 = int64(uint64(b[0]) << 56)
	// 如果不强制转换为long，那么默认会当作int，导致最高32位丢失
	l = int64(uint64(l) & uint64(0xFF00000000000000))
	l |= int64(uint64(b[1]) << 48) & 0xFF000000000000
	l |= int64(uint64(b[2]) << 40) & 0xFF0000000000
	l |= int64(uint64(b[3]) << 32) & 0xFF00000000
	l |= int64(uint64(b[4]) << 24) & 0xFF000000
	l |= int64(uint64(b[5]) << 16) & 0xFF0000
	l |= int64(uint64(b[6]) << 8) & 0xFF00
	l |= int64(uint64(b[7])) & 0xFF
	return l
}

func BytesHighFirstToInt64(bb []byte) int64 {
	var b []byte = make([]byte, 8)
	copy(b, bb)
	var l int64 = int64(uint64(b[0]) << 56)
	l = int64(uint64(l) & uint64(0xFF00000000000000))
	l |= int64(uint64(b[1]) << 48) & 0xFF000000000000
	l |= int64(uint64(b[2]) << 40) & 0xFF0000000000
	l |= int64(uint64(b[3]) << 32) & 0xFF00000000
	l |= int64(uint64(b[4]) << 24) & 0xFF000000
	l |= int64(uint64(b[5]) << 16) & 0xFF0000
	l |= int64(uint64(b[6]) << 8) & 0xFF00
	l |= int64(uint64(b[7])) & 0xFF
	return l
}

/**
 * 将一个char字符转换位字节数组（2个字节），b[0]存储高位字符，大端
 *
 * @param c 字符（java char 2个字节）
 * @return 代表字符的字节数组
 */
func CharToBytes(c Char) []byte {
	var b []byte = make([]byte, 8)
	b[0] = byte(uint(c) >> 8)
	b[1] = byte(c)
	return b
}

/**
 * 将一个双精度浮点数转换位字节数组（8个字节），b[0]存储高位字符，大端
 *
 * @param d 双精度浮点数
 * @return 代表双精度浮点数的字节数组
 */
func Float64ToBytes(d float64) []byte {
	return Int64ToBytes(gconv.Int64(d))
}

/**
 * 将一个整数转换位字节数组(4个字节)，b[0]存储高位字符，大端
 *
 * @param i 整数
 * @return 代表整数的字节数组
 */
func IntToBytes(i int) []byte {
	var b = make([]byte, 4)
	var v uint = uint(i)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b
}

/**
 * 将一个长整数转换位字节数组(8个字节)，b[0]存储高位字符，大端
 *
 * @param l 长整数
 * @return 代表长整数的字节数组
 */
func Int64ToBytes(l int64) []byte {
	var b []byte = make([]byte, 8)
	var v uint64 = uint64(l)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

func ConvertTwoCharToInt(high Char, low Char) int {
	var result int = int(high) << 16
	result |= int(low)
	return result
}

func ConvertIntToTwoChar(n int) []Char {
	var result []Char = make([]Char, 2)
	result[0] = Char(uint(n) >> 16)
	result[1] = Char(0x0000FFFF & uint(n))
	return result
}
