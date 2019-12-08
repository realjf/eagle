package trie

type MatchToken struct {
	Token
	emit Emit
}

func NewMatchToken(fragment string, emit Emit) *MatchToken {
	return &MatchToken{
		emit:  emit,
		Token: Token{fragment: fragment},
	}
}

func (mt *MatchToken) IsMatch() bool {
	return true
}

func (mt *MatchToken) GetEmit() Emit {
	return mt.emit
}
