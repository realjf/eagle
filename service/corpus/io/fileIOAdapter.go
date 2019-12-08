package io

import "io"

/**
 * 基于普通文件系统的IO适配器
 *
 * @author hankcs
 */
type FileIOAdapter struct {
	IIOAdapter
}

func (f *FileIOAdapter) Open(path string) (io.Reader, error) {
	return nil, nil
}

func (f *FileIOAdapter) Create(path string) (io.Writer, error) {
	return nil, nil
}


