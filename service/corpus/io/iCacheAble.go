package io

import ByteArray "eagle/service/corpus/io/byteArray"

/**
 * 可写入或读取二进制
 * @author hankcs
 */
type ICacheAble interface {
	Save(stream DataOutputStream) error // 写入
	Load(array ByteArray.ByteArray) bool // 载入
}


