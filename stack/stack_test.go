package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	if s.Length() != 0 {
		t.Errorf("expected length 0, got %d", s.Length())
	}
}

func TestWrapStack(t *testing.T) {
	s := Wrap([]string{"a", "b", "c"})
	if s.Length() != 3 {
		t.Errorf("expected length 3, got %d", s.Length())
	}
}

func TestPush(t *testing.T) {
	s := NewStack[string]()

	length := s.Push("hello")
	if length != 1 {
		t.Errorf("expected length 1 after first push, got %d", length)
	}

	length = s.Push("world")
	if length != 2 {
		t.Errorf("expected length 2 after second push, got %d", length)
	}

	if s.Length() != 2 {
		t.Errorf("Length mismatch: expected 2, got %d", s.Length())
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int]()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	val, length := s.Pop()
	if val != 30 || length != 2 {
		t.Errorf("expected (30, 2), got (%d, %d)", val, length)
	}

	val, length = s.Pop()
	if val != 20 || length != 1 {
		t.Errorf("expected (20, 1), got (%d, %d)", val, length)
	}

	val, length = s.Pop()
	if val != 10 || length != 0 {
		t.Errorf("expected (10, 0), got (%d, %d)", val, length)
	}
}

func TestPopEmpty(t *testing.T) {
	s := NewStack[float64]()

	val, length := s.Pop()
	var zero float64

	if length != 0 {
		t.Errorf("expected length 0, got %d", length)
	}
	if !reflect.DeepEqual(val, zero) {
		t.Errorf("expected zero value, got %v", val)
	}
}
