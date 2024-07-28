package ds

import (
	"bytes"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		testName           string
		numsToPush         []int
		numsToPop          int
		expectedQueueState *Queue[int]
	}{
		{
			testName:   "empty queue",
			numsToPush: []int{},
			numsToPop:  0,
			expectedQueueState: &Queue[int]{
				front: 0,
				cap:   10,
				rear:  9,
				size:  0,
				data:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			testName:   "only push",
			numsToPush: []int{1, 2, 3, 4, 5},
			numsToPop:  0,
			expectedQueueState: &Queue[int]{
				front: 0,
				rear:  4,
				size:  5,
				cap:   10,
				data:  []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0},
			},
		},
		{
			testName:   "push and pop",
			numsToPush: []int{1, 2, 3, 4, 5},
			numsToPop:  5,
			expectedQueueState: &Queue[int]{
				front: 5,
				rear:  4,
				size:  0,
				cap:   10,
				data:  []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			actualQ := NewDefaultQueue[int]()
			for _, e := range test.numsToPush {
				err := actualQ.Enqueue(e)
				if err != nil {
					tt.Errorf("error: expected nil, but got %s", err.Error())
				}
			}
			for range test.numsToPop {
				_, err := actualQ.Dequeue()
				if err != nil {
					tt.Errorf("error: expected nil, but got %s", err.Error())
				}
			}
			if !reflect.DeepEqual(test.expectedQueueState, actualQ) {
				tt.Errorf("expected %+v, but got %+v", test.expectedQueueState, actualQ)
			}
		})
	}
}

func TestQueuePrint(t *testing.T) {
	q := NewDefaultQueue[int]()
	expectedQPrint := "Q: 0 1 2 3 4 5 6 7 8 9"
	for i := range(10) {
		err := q.Enqueue(i)
		if err != nil {
			t.Errorf("error: expected nil, but got %s", err.Error())
		}
	}
	var qBuffer bytes.Buffer
	q.Print(&qBuffer)
	if expectedQPrint != qBuffer.String() {
		t.Errorf("expected %s, but got %s", expectedQPrint, qBuffer.String())
	}
}
