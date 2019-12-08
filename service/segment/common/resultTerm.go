package common

import (
	"fmt"
)

// 一个通用的Term
type ResultTerm struct {
	Word string
	Label interface{}
	Offset int
}

func NewResultTerm(word string, label interface{}, offset int) *ResultTerm {
	return &ResultTerm{
		Word:word,
		Label:label,
		Offset:offset,
	}
}

func (rt *ResultTerm) ToString() string {
	return fmt.Sprintf("%s/%v", rt.Word, rt.Label)
}


