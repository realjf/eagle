package trie

type FragmentToken struct {
	Token
}

func NewFragmentToken(fragment string) *FragmentToken {
	return &FragmentToken{
		Token{
			fragment: fragment,
		},
	}
}

func (ft *FragmentToken) IsMatch() bool {
	return false
}

func (ft *FragmentToken) GetEmit() Emit {
	return Emit{}
}
