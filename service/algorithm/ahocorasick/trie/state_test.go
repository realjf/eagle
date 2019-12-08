package trie

import "testing"

func TestNewState2(t *testing.T) {
	state := NewState2(1)
	t.Fatal(state.GetDepth())
}
