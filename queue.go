package ds

import (
	"errors"
	"fmt"
	"io"
)

const (
	ERROR_EMPTY_QUEUE = "queue is empty"
)

type Queue[T comparable] struct {
	size, cap, front, rear int 
	data []T
}

func NewQueue[T comparable](capacity int) *Queue[T] {
	return &Queue[T]{
		size: 0,
		cap: capacity,
		front: 0,
		rear: capacity - 1,
		data: make([]T, capacity),
	}
}

func NewDefaultQueue[T comparable]() *Queue[T] {
	return NewQueue[T](10)
}

func (q *Queue[T]) Enqueue(data T) error {
	if q.IsFull() {
		return errors.New(ERROR_OVERFLOW)
	}
	q.rear = (q.rear + 1) % q.cap
	q.data[q.rear] = data
	q.size++
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var data T	
	if q.IsEmpty() {
		return data, errors.New(ERROR_UNDERFLOW)
	}
	data = q.data[q.front]
	q.front = (q.front + 1) % q.cap
	q.size--
	return data, nil
}

func (q *Queue[T]) IsFull() bool {
	return q.cap == q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Print(w io.Writer) {
	w.Write([]byte("Q:"))
	for _, e := range q.data {
		w.Write([]byte(fmt.Sprintf(" %v", e)))
	}
}
