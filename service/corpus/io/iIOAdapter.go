package io

import "io"

/**
 * IO适配器接口<br>
 * 实现该接口以移植HanLP到不同的平台
 *
 * @author hankcs
 */
type IIOAdapter interface {
	Open(path string) (io.Reader, error) // 打开一个文件以供读取
	Create(path string) (io.Writer, error) // 创建一个新文件以供输出
}
