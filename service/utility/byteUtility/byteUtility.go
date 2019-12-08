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
func BytesHighFirstToInt(bytes []byte, start int) int {
	var num int = gconv.Int(bytes[start+3] & 0xFF)
	num |= gconv.Int(bytes[start+2] << 8 ) & 0xFF00
	num |= gconv.Int(bytes[start+1] << 16 ) & 0xFF0000
	num |= gconv.Int(bytes[start] << 24 ) & 0xFF000000
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


func BytesHighFirstToFloat64(bytes []byte, start int) float64 {
	var l int64 = int64(bytes[start] << 56) & 0xFF00000000000000
	// 如果不强制转换为long，那么默认会当作int，导致最高32位丢失
	l |= int64(bytes[start + 1] << 48) & 0xFF000000000000
	l |= int64(bytes[start + 2] << 40) & 0xFF0000000000
	l |= int64(bytes[start + 3] << 32) & 0xFF00000000
	l |= int64(bytes[start + 4] << 24) & 0xFF000000
	l |= int64(bytes[start + 5] << 16) & 0xFF0000
	l |= int64(bytes[start + 6] << 8) & 0xFF00
	l |= int64(bytes[start + 7]) & 0xFF

	return gconv.Float64(l)
}


func BytesHighFirstToFloat(bytes []byte, start int) float32 {
	var l int = BytesHighFirstToInt(bytes, start)
	return gconv.Float32(l)
}

