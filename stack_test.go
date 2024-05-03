package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	testCases := []struct{
		testName string
		size int
		pushes []int
		expected int
		expectedError error
	}{
		{
			testName: "default stack",
			size: 0,
			pushes: []int{2,5,3,4},
			expected: 4,
			expectedError: nil,	
		},
		{
			testName: "only pushes",
			size: 5,
			pushes: []int{2,5,3,4},
			expected: 4,
			expectedError: nil,
		},
		{
			testName: "empty stack",
			size: 5,
			pushes: []int{},
			expected: 0,
			expectedError: fmt.Errorf(ERROR_EMPTY_STACK),
		},
	}
	for _, test := range testCases {
		t.Run(test.testName, func(tt *testing.T) {
			var st *Stack	
			if test.size == 0 {
				st = NewDefaultStack() 
			} else {
				st = NewStack(test.size)
			}
			for _, e := range test.pushes {
				st.Push(e)
			}
			result, err := st.Top()
			assert.Equal(tt, test.expected, result)
			assert.Equal(tt, test.expectedError, err)
		})
	}
}
