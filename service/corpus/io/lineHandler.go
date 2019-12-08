package io

import (
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

type LineHandler struct {
	Delimiter string
}

func NewLineHandler(delimiter string) *LineHandler {
	if delimiter == "" {
		delimiter = "\t"
	}
	return &LineHandler{Delimiter:delimiter}
}

func (l *LineHandler) Handle(line string) error {
	tokenList := glist.New(true)
	start := 0
	var end int
	for end = gstr.Pos(line, l.Delimiter, start); end != -1; {
		tokenList.PushBack(gstr.SubStr(line, start, end))
		start = end + 1
	}
	tokenList.PushBack(gstr.SubStr(line, start, len(line)))
	return l.Handle2(gconv.SliceStr(tokenList.FrontAll()))
}

func (l *LineHandler) Done() error {
	return nil
}

func (l *LineHandler) Handle2(params []string) error {
	return nil
}

