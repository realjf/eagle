package io

type FileInputStream struct {
	filePath string
}

func NewFileInputStream(path string) *FileInputStream {
	return &FileInputStream{
		filePath:path,
	}
}



