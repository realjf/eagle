package ByteArrayStream

import (
	"eagle/service/common"
	"eagle/service/corpus/io"
	. "eagle/service/corpus/io/ByteArray"
	"eagle/service/config"
	"eagle/utils"
)

type ByteArrayStream struct {
	ByteArray

	BufferSize int
}

func NewByteArrayStream(bytes []byte, bufferSize int) *ByteArrayStream {
	bas := &ByteArrayStream{
		BufferSize: bufferSize,
		ByteArray: ByteArray{
			Bytes:bytes,
		},
	}

	return bas
}

func CreateByteArrayStream(path string) *ByteArrayStream {
	if config.GConfig.IOAdapter == nil {
		return
	}
	is, err := config.GConfig.IOAdapter.Open(path)
	if err != nil {
		utils.Logger.Warning("打开失败： " + path)
		return nil
	}
	if is.(type) == io.FileInputStream {
		return
	}
}

func (bas *ByteArrayStream) NextInt() int {
	bas.EnsureAvailableBytes(4)
	return bas.ByteArray.NextInt()
}

func (bas *ByteArrayStream) NextChar() common.Char {
	bas.EnsureAvailableBytes(2)
	return bas.ByteArray.NextChar()
}

func (bas *ByteArrayStream) NextFloat64() float64 {
	bas.EnsureAvailableBytes(8)
	return bas.ByteArray.NextFloat64()
}

func (bas *ByteArrayStream) NextByte() byte {
	bas.EnsureAvailableBytes(1)
	return bas.ByteArray.NextByte()
}

func (bas *ByteArrayStream) NextFloat() float32 {
	bas.EnsureAvailableBytes(4)
	return bas.ByteArray.NextFloat()
}

func (bas *ByteArrayStream) EnsureAvailableBytes(size int) {

}

