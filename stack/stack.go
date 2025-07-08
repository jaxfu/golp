package stack

import "slices"

type Stack[T any] struct {
	slice []T
}

// create Stack and slice of prealloc capacity
func NewStack[T any]() Stack[T] {
	slice := make([]T, 0, 16)
	return Stack[T]{
		slice: slice,
	}
}

// wrap premade slice, use as List
func Wrap[T any](slice []T) Stack[T] {
	return Stack[T]{
		slice: slice,
	}
}

// returns length after push
func (s *Stack[T]) Curr() (T, uint) {
	ln := uint(len(s.slice))
	if ln <= 0 {
		return *new(T), 0
	}

	return s.slice[ln-1], ln
}

// returns length after push
func (s *Stack[T]) Push(el T) uint {
	s.slice = append(s.slice, el)

	return uint(len(s.slice))
}

// returns popped element, resulting length
func (s *Stack[T]) Pop() (T, uint) {
	ln := len(s.slice)
	if ln <= 0 {
		return *new(T), 0
	}

	el := s.slice[ln-1]
	s.slice = slices.Delete(s.slice, ln-1, ln)

	return el, uint(len(s.slice))
}

func (s *Stack[T]) Length() uint {
	return uint(len(s.slice))
}
