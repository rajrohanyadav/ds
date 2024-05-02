package ds

import "fmt"

type Stack struct {
	top int
	size int
	data []int
}

func NewDefaultStack() *Stack {
	return &Stack{
		top: 0,
		size: 10,
		data: make([]int, 10),
	}
}

func (s Stack) Top() int {
	return s.data[s.top]
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
