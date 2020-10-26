package ByteArrayFileStream

import (
	"eagle/service/corpus/io"
	"eagle/service/corpus/io/ByteArrayStream"
	"os"
)

type ByteArrayFileStream struct {
	ByteArrayStream.ByteArrayStream
	FileChannel *os.File
}

func CreateByteArrayFileStream(path string) *ByteArrayFileStream {
	return nil
}

func CreateByteArrayFileStream2(fileInputStream io.FileInputStream) *ByteArrayFileStream {
	return nil
}

func (b *ByteArrayFileStream) HasMore() bool {
	return  b.Offset < b.BufferSize || b.FileChannel != nil
}

