package trie

import "fmt"

type IToken interface {
	GetFragment() string
	IsMatch() bool
	GetEmit() Emit
	ToString() string
}

// 一个片段
type Token struct {
	fragment string // 对应的片段
}

func NewToken(fragment string) *Token {
	return &Token{
		fragment: fragment,
	}
}

func (t *Token) GetFragment() string {
	return t.fragment
}

func (t *Token) IsMatch() bool {
	return false
}

func (t *Token) GetEmit() Emit {
	return Emit{}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s/%v", t.fragment, t.IsMatch())
}
