package ahocorasick

import (
	. "eagle/service/common"
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState2(1)
	t.Log(state.ToString())
	state.AddEmit(2)
	t.Log(state.ToString())
	state.AddState(Char('b'))
	t.Log(state.AddState(Char('c')))
	t.Log(state.ToString())
	state.NextState(Char('a'), true)
	t.Log(state.ToString())
	t.Log(state.GetStates())
	t.Log(state.GetDepth())
	t.Log(state.GetTransitions())
	t.Log(state.GetIndex())
	t.Log(state.GetSuccess())
	t.Log(state.GetLargestValueId())
	t.Fatal(state)
}
