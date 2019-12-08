package io

import (
	"bytes"
	"errors"
	"eagle/utils"
	"github.com/gogf/gf/util/gconv"
	"io"
	"io/ioutil"
)

// 二进制数据输入流
type IDataInputStream interface {
	io.ByteReader
	io.Reader
}

type DataInputStream struct {
	filePath string
	fileByte []byte
	buffer *bytes.Buffer
	reader *bytes.Reader
}

func NewDataInputStream(path string) *DataInputStream {
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger.Warning("open file " + path + " error: " + err.Error())
		return nil
	}

	reader := bytes.NewReader(fileByte)

	return &DataInputStream {
		filePath: path,
		fileByte: fileByte,
		buffer: bytes.NewBuffer(fileByte),
		reader: reader,
	}
}

func (dis *DataInputStream) ReadByte() (byte, error) {
	if dis.reader == nil {
		return 0, errors.New("no file handler")
	}
	char, err := dis.reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return gconv.Byte(char), nil
}

func (dis *DataInputStream) ReadInt() int {
	ch1, err := dis.ReadByte()
	if err != nil {
		panic("read byte error: " + err.Error())
	}
	ch2, err := dis.ReadByte()
	if err != nil {
		panic("read byte error: " + err.Error())
	}
	ch3, err := dis.ReadByte()
	if err != nil {
		panic("read byte error: " + err.Error())
	}
	ch4, err := dis.ReadByte()
	if err != nil {
		panic("read byte error: " + err.Error())
	}
	return gconv.Int((ch1 << 24) + (ch2 << 16) + (ch3 << 8) + (ch4 << 0))
}

func (dis *DataInputStream) Reset() error {
	if dis.reader == nil {
		return errors.New("no file handler")
	}

	dis.reader.Reset(dis.fileByte)
	return nil
}

func (dis *DataInputStream) Read(p []byte) (int, error) {
	if dis.reader == nil {
		return 0, errors.New("no file handler")
	}
	return dis.reader.Read(p)
}
