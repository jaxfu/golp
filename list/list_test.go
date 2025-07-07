package list

import (
	"testing"
)

func TestNewList(t *testing.T) {
	var prealloc uint = 16
	l := NewList[int](prealloc)
	if len(l.slice) != 0 {
		t.Errorf("expected length 0, got %d", len(l.slice))
	}
	if cap(l.slice) != int(prealloc) {
		t.Errorf("expected capacity %d, got %d", prealloc, cap(l.slice))
	}
	if l.cursor != 0 {
		t.Errorf("expected cursor 0, got %d", l.cursor)
	}
}

func TestWrap(t *testing.T) {
	s := []int{10, 20, 30}
	l := Wrap(s)
	if len(l.slice) != 3 {
		t.Errorf("expected length 3, got %d", len(l.slice))
	}
	if l.Curr() != 10 {
		t.Errorf("expected current element 10, got %d", l.Curr())
	}
}

func TestMove(t *testing.T) {
	l := Wrap([]int{1, 2, 3, 4, 5})

	val := l.Move(2)
	if val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
	if l.Cursor() != 2 {
		t.Errorf("expected cursor 2, got %d", l.Cursor())
	}

	val = l.Move(-1)
	if val != 2 {
		t.Errorf("expected 2, got %d", val)
	}

	val = l.Move(-10)
	if val != 1 {
		t.Errorf("expected 1, got %d", val)
	}
	if l.Cursor() != 0 {
		t.Errorf("expected cursor 0, got %d", l.Cursor())
	}

	val = l.Move(10)
	if val != 5 {
		t.Errorf("expected 5, got %d", val)
	}
	if l.Cursor() != 4 {
		t.Errorf("expected cursor 4, got %d", l.Cursor())
	}
}

func TestLook(t *testing.T) {
	l := Wrap([]int{100, 200, 300})

	if v := l.Look(0); v != 100 {
		t.Errorf("expected 100, got %d", v)
	}
	if v := l.Look(1); v != 200 {
		t.Errorf("expected 200, got %d", v)
	}
	if v := l.Look(5); v != 300 {
		t.Errorf("expected 300 (overflow), got %d", v)
	}
	if v := l.Look(-1); v != 100 {
		t.Errorf("expected 100 (underflow), got %d", v)
	}
}
