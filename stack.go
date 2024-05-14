package ds

import "fmt"

const (
	ERROR_EMPTY_STACK = "stack is empty"
	ERROR_OVERFLOW = "overflow"
	ERROR_UNDERFLOW = "underflow"
)

type Stack[T any] struct {
	top int 
	size int
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		top: -1,
		size: size,
		data: make([]T, size),
	}
}

func NewDefaultStack[T any]() *Stack[T] {
	return NewStack[T](10)
}

func (s Stack[T]) Top() (T, error) {
	var ret T
	if s.top < 0 {
		return ret, fmt.Errorf(ERROR_EMPTY_STACK)
	}
	ret = s.data[s.top]
	return ret, nil
}

func (s *Stack[T]) Push(val T) error {
	if s.top + 1 == s.size {
		return fmt.Errorf("Overflow")
	}
	s.top++
	s.data[s.top] = val 
	return nil
}

func (s *Stack[T]) Pop() error {
	if s.top == 0 {
		return fmt.Errorf("Underflow")
	}
	s.top--
	return nil
}
