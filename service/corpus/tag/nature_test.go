package tag

import "testing"

func TestNewNature(t *testing.T) {
	nature := NewNature("g")
	t.Fatalf("%v", nature)
}
