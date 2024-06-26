package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Remove duplication from test code

func TestStackIntegerT(t *testing.T) {
	testCases := []struct{
		testName string
		size int
		pushes []int
		expected int
		expectedIsEmpty bool
		expectedError error
	}{
		{
			testName: "default stack",
			size: 0,
			pushes: []int{2,5,3,4},
			expected: 4,
			expectedIsEmpty: false,
			expectedError: nil,	
		},
		{
			testName: "only pushes",
			size: 5,
			pushes: []int{2,5,3,4},
			expected: 4,
			expectedIsEmpty: false,
			expectedError: nil,
		},
		{
			testName: "empty stack",
			size: 5,
			pushes: []int{},
			expected: 0,
			expectedIsEmpty: true,
			expectedError: fmt.Errorf(ERROR_EMPTY_STACK),
		},
	}
	for _, test := range testCases {
		t.Run(test.testName, func(tt *testing.T) {
			var st *Stack[int]	
			if test.size == 0 {
				st = NewDefaultStack[int]() 
			} else {
				st = NewStack[int](test.size)
			}
			for _, e := range test.pushes {
				st.Push(e)
			}
			result, err := st.Top()
			assert.Equal(tt, test.expected, result)
			assert.Equal(tt, test.expectedError, err)
			assert.Equal(tt, test.expectedIsEmpty, st.IsEmpty())
		})
	}
}

func TestStackFloatT(t *testing.T) {
	testCases := []struct{
		testName string
		size int
		pushes []float32
		expected float32
		expectedIsEmpty bool
		expectedError error
	}{
		{
			testName: "default stack",
			size: 0,
			pushes: []float32{2.2,5.23,3.234,4.342},
			expected: 4.342,
			expectedIsEmpty: false,
			expectedError: nil,	
		},
		{
			testName: "only pushes",
			size: 5,
			pushes: []float32{2,5,3,4},
			expected: 4,
			expectedIsEmpty: false,
			expectedError: nil,
		},
		{
			testName: "empty stack",
			size: 5,
			pushes: []float32{},
			expected: 0,
			expectedIsEmpty: true,
			expectedError: fmt.Errorf(ERROR_EMPTY_STACK),
		},
	}
	for _, test := range testCases {
		t.Run(test.testName, func(tt *testing.T) {
			var st *Stack[float32]	
			if test.size == 0 {
				st = NewDefaultStack[float32]() 
			} else {
				st = NewStack[float32](test.size)
			}
			for _, e := range test.pushes {
				st.Push(e)
			}
			result, err := st.Top()
			assert.Equal(tt, test.expected, result)
			assert.Equal(tt, test.expectedError, err)
			assert.Equal(tt, test.expectedIsEmpty, st.IsEmpty())
		})
	}
}

func TestStackStringT(t *testing.T) {
	testCases := []struct{
		testName string
		size int
		pushes []string
		expected string 
		expectedIsEmpty bool
		expectedError error
	}{
		{
			testName: "default stack",
			size: 0,
			pushes: []string{"hello", "world"},
			expected: "world",
			expectedIsEmpty: false,
			expectedError: nil,	
		},
		{
			testName: "only pushes",
			size: 5,
			pushes: []string{"2", "5", "5", "4"},
			expected: "4",
			expectedIsEmpty: false,
			expectedError: nil,
		},
		{
			testName: "empty stack",
			size: 5,
			pushes: []string{},
			expected: "",
			expectedIsEmpty: true,
			expectedError: fmt.Errorf(ERROR_EMPTY_STACK),
		},
	}
	for _, test := range testCases {
		t.Run(test.testName, func(tt *testing.T) {
			var st *Stack[string]	
			if test.size == 0 {
				st = NewDefaultStack[string]() 
			} else {
				st = NewStack[string](test.size)
			}
			for _, e := range test.pushes {
				st.Push(e)
			}
			result, err := st.Top()
			assert.Equal(tt, test.expected, result)
			assert.Equal(tt, test.expectedError, err)
			assert.Equal(tt, test.expectedIsEmpty, st.IsEmpty())
		})
	}
}

func TestStackInterfaceT(t *testing.T) {
	testCases := []struct{
		testName string
		size int
		pushes []interface{}
		expected interface{}
		expectedIsEmpty bool 
		expectedError error
	}{
		{
			testName: "default stack",
			size: 0,
			pushes: []interface{}{"hello", "world"},
			expectedIsEmpty: false,
			expected: "world",
			expectedError: nil,	
		},
		{
			testName: "only pushes",
			size: 5,
			pushes: []interface{}{2, "5", "5", 4},
			expected: 4,
			expectedIsEmpty: false,
			expectedError: nil,
		},
		{
			testName: "empty stack",
			size: 5,
			pushes: []interface{}{},
			expected: nil,
			expectedIsEmpty: true,
			expectedError: fmt.Errorf(ERROR_EMPTY_STACK),
		},
	}
	for _, test := range testCases {
		t.Run(test.testName, func(tt *testing.T) {
			var st *Stack[interface{}]	
			if test.size == 0 {
				st = NewDefaultStack[interface{}]() 
			} else {
				st = NewStack[interface{}](test.size)
			}
			for _, e := range test.pushes {
				st.Push(e)
			}
			result, err := st.Top()
			assert.Equal(tt, test.expected, result)
			assert.Equal(tt, test.expectedError, err)
			assert.Equal(tt, test.expectedIsEmpty, st.IsEmpty())
		})
	}
}
