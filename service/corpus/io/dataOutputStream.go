package io

import (
	"errors"
	"eagle/utils"
	"github.com/gogf/gf/util/gconv"
	"fmt"
	"io"
	"os"
)

// 二进制数据输出流
type IDataOutputStream interface {
	io.ByteWriter
}

type DataOutputStream struct {
	file *os.File
}

func NewDataOutputStream(path string) *DataOutputStream {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		utils.Logger.Warning("open file " + path + " error: " + err.Error())
		return nil
	}

	return &DataOutputStream {
		file: file,
	}
}

func (dos *DataOutputStream) WriteByte(c byte) error {
	if dos.file == nil {
		return errors.New("no file handler")
	}

	n, err := dos.file.Write(gconv.Bytes(c))
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("write 0 byte")
	}
	return nil
}

func (dos *DataOutputStream) WriteString(s string) (n int, err error) {
	if dos.file == nil {
		return 0, errors.New("no file handler")
	}

	n, err = dos.file.Write(gconv.Bytes(s))
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return n, errors.New("write 0 byte")
	}
	return n,nil
}

func (dos *DataOutputStream) WriteInt(v int) (n int, err error) {
	if dos.file == nil {
		return 0, errors.New("no file handler")
	}
	data := []byte{}
	data = append(data, gconv.Byte(v >> 24 & 255))
	data = append(data, gconv.Byte(v >> 16 & 255))
	data = append(data, gconv.Byte(v >> 8 & 255))
	data = append(data, gconv.Byte(v >> 0 & 255))
	fmt.Println(data)
	n, err = dos.file.Write(data)
	if err != nil {
		return n, err
	}
	if n == 0 {
		return n, errors.New("write 0 byte")
	}
	return n,nil
}


func (dos *DataOutputStream) WriteChar(v int) (n int, err error) {
	n, err = dos.WriteInt((v >> 8) & 0xFF)
	if err != nil {
		return
	}
	n1, err := dos.WriteInt((v >> 0) & 0xFF)
	if err != nil {
		return
	}
	n += n1
	return
}


func (dos *DataOutputStream) Close() error {
	return dos.file.Close()
}