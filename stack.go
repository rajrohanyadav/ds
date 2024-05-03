package ds

import "fmt"

const (
	ERROR_EMPTY_STACK = "stack is empty"
	ERROR_OVERFLOW = "overflow"
	ERROR_UNDERFLOW = "underflow"
)

type Stack struct {
	top int
	size int
	data []int
}

func NewStack(size int) *Stack {
	return &Stack{
		top: -1,
		size: size,
		data: make([]int, size),
	}
}

func NewDefaultStack() *Stack {
	return NewStack(10)
}

func (s Stack) Top() (int, error) {
	if s.top < 0 {
		return 0, fmt.Errorf(ERROR_EMPTY_STACK)
	}
	return s.data[s.top], nil
}

func (s *Stack) Push(val int) error {
	if s.top + 1 == s.size {
		return fmt.Errorf("Overflow")
	}
	s.top++
	s.data[s.top] = val 
	return nil
}

func (s *Stack) Pop() error {
	if s.top == 0 {
		return fmt.Errorf("Underflow")
	}
	s.top--
	return nil
}
