package ds

import (
	"bytes"
	"testing"
)

func TestBTreeTraversals(t *testing.T) {
	tests := []struct{
		testName string
		inputArr []int
		expectedInOrder string
		expectedPreOrder string
		expectedPostOrder string
	}{
		{
			testName: "empty tree",
			inputArr: []int{},
			expectedInOrder: "",
			expectedPreOrder: "",
			expectedPostOrder: "",
		},
		{
			testName: "non-empty tree",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			expectedInOrder: " 🌳1 🌳1 🌳2 🌳3 🌳5 🌳6 🌳21 🌳23",
			expectedPreOrder: " 🌳1 🌳1 🌳5 🌳2 🌳3 🌳23 🌳21 🌳6",
			expectedPostOrder: " 🌳1 🌳3 🌳2 🌳6 🌳21 🌳23 🌳5 🌳1",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			var inorder, preorder, postorder bytes.Buffer
			bTree := NewBTree()
			bTree.ParseFromArray(test.inputArr)
			bTree.root.PrintInOrder(&inorder)
			bTree.root.PrintPreOrder(&preorder)
			bTree.root.PrintPostOrder(&postorder)
			if inorder.String() != test.expectedInOrder {
				tt.Errorf("inorder: got %s but expected %s", inorder.String(), test.expectedInOrder)
			}
			if preorder.String() != test.expectedPreOrder {
				tt.Errorf("preorder: got %s but expected %s", preorder.String(), test.expectedPreOrder)
			}
			if postorder.String() != test.expectedPostOrder {
				tt.Errorf("postorder: got %s but expected %s", postorder.String(), test.expectedPostOrder)
			}	
		})
	}
}

func TestBTreeDelete(t *testing.T) {
	tests := []struct{
		testName string
		inputArr []int
		deleteArr []int
		expectedInOrder string
	}{
		{
			testName: "empty tree",
			inputArr: []int{},
			deleteArr: []int{},
			expectedInOrder: "",
		},
		{
			testName: "no delete",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			deleteArr: []int{},
			expectedInOrder: " 🌳1 🌳1 🌳2 🌳3 🌳5 🌳6 🌳21 🌳23",
		},
		{
			testName: "leaf delete",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			deleteArr: []int{6},
			expectedInOrder: " 🌳1 🌳1 🌳2 🌳3 🌳5 🌳21 🌳23",
		},
		{
			testName: "single child delete",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			deleteArr: []int{21},
			expectedInOrder: " 🌳1 🌳1 🌳2 🌳3 🌳5 🌳6 🌳23",
		},
		{
			testName: "two children delete",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			deleteArr: []int{5},
			expectedInOrder: " 🌳1 🌳1 🌳2 🌳3 🌳6 🌳21 🌳23",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			var inorder bytes.Buffer
			bTree := NewBTree()
			bTree.ParseFromArray(test.inputArr)
			for _, e := range test.deleteArr {
				err := bTree.Delete(e)
				if err != nil {
					tt.Errorf("expected nil error, but got %s", err.Error())
				}
			}
			bTree.root.PrintInOrder(&inorder)
			if inorder.String() != test.expectedInOrder {
				tt.Errorf("inorder: got %s but expected %s", inorder.String(), test.expectedInOrder)
			}
		})
	}
}

func TestBTreeMaxMin(t *testing.T) {
	tests := []struct{
		testName string
		inputArr []int
		expectedMax int
		expectedMin int
	}{
		{
			testName: "empty tree",
			inputArr: []int{},
			expectedMax: 0,
			expectedMin: 0,
		},
		{
			testName: "tree with values",
			inputArr: []int{1, 5, 23, 2, 21, 1, 6, 3},
			expectedMax: 23,
			expectedMin: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			bTree := NewBTree()
			bTree.ParseFromArray(test.inputArr)
			actualMax := bTree.Max()
			actualMin := bTree.Min()
			if actualMax != test.expectedMax {
				tt.Errorf("max: expected %d, but got %d", test.expectedMax, actualMax)
			}
			if actualMin != test.expectedMin {
				tt.Errorf("min: expected %d, but got %d", test.expectedMin, actualMin)
			}
		})
	}
}
