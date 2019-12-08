package trie

import "testing"

func TestNewEmit(t *testing.T) {
	emit := NewEmit(1, 2, "hello")
	t.Fatal(emit.ToString(), emit.GetKeyword(), emit.GetStart(), emit.GetEnd(), emit.HashCode())
}
