package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	s.Push(5)

	val, flag := s.Pop()
	if val != 5 || flag != true {
		t.Log("TestStack s.Pop failed")
		t.Fail()
	}
}
