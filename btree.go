package ds

import (
	"errors"
	"fmt"
	"io"
)

// TODO: revisit errors and document

const (
	ERROR_ITEM_NOT_FOUND = "data not found"
	ERROR_OPERATION_NOT_ALLOWED_ON_NIL = "operation not allowed on nil"
	ERROR_TREE_IS_EMPTY = "tree is empty"
)

type BTree struct {
	root *BTreeNode
}

type BTreeNode struct {
	val   int
	left  *BTreeNode
	right *BTreeNode
}

func NewBTree() *BTree {
	return &BTree{
		root: nil,
	}
}

func NewBTreeNode(data int) *BTreeNode {
	return &BTreeNode{
		val:   data,
		left:  nil,
		right: nil,
	}
}

func (t *BTree) Insert(data int) {
	newNode := NewBTreeNode(data)
	if t.root == nil {
		t.root = newNode
		return
	}
	curr := t.root
	var parent *BTreeNode
	for curr != nil {
		parent = curr
		if newNode.val <= curr.val {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	if newNode.val <= parent.val {
		parent.left = newNode
	} else {	
		parent.right = newNode
	}
}

func (t *BTree) Delete(data int) error {
	if t == nil {
		return errors.New(ERROR_TREE_IS_EMPTY)
	}
	fakeParent := &BTreeNode{right: t.root}
	err := t.root.Delete(data, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.right == nil {
		t.root = nil
	}
	return nil
}

func (n *BTreeNode) Delete(data int, parent *BTreeNode) error {
	if n == nil {
		return errors.New(ERROR_ITEM_NOT_FOUND)
	}
	switch {
	case data < n.val:
		return n.left.Delete(data, n)
	case data > n.val:
		return n.right.Delete(data, n)
	default:
		if n.left == nil && n.right == nil {
			n.ReplaceNode(parent, nil)
			return nil
		}

		if n.left == nil {
			n.ReplaceNode(parent, n.right)
			return nil
		}

		if n.right == nil {
			n.ReplaceNode(parent, n.left)
			return nil
		}
		
		replacement, replacementParent := n.left.MaxNode(n)
		n.val = replacement.val
		
		return replacement.Delete(replacement.val, replacementParent)
	}
}

func (n *BTreeNode) ReplaceNode(parent, replacement *BTreeNode) error {
	if n == nil {
		return errors.New(ERROR_OPERATION_NOT_ALLOWED_ON_NIL)
	}

	if n == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}

func(n *BTreeNode) MinNode(parent *BTreeNode) (*BTreeNode, *BTreeNode) {
	if n == nil {
		return nil, parent
	}	
	for n.left != nil {
		parent = n 
		n = n.left
	}
	return n, parent 
}

func (n *BTreeNode) MaxNode(parent *BTreeNode) (*BTreeNode, *BTreeNode) {
	if n == nil {
		return nil, parent
	}
	for n.right != nil {
		parent = n 
		n = n.right
	}
	return n, parent
}

func (t *BTree) Min() int {
	if t.root == nil {
		return 0
	}
	curr := t.root
	for curr.left != nil {
		curr = curr.left
	}
	return curr.val
}

func (t *BTree) Max() int {
	if t.root == nil {
		return 0
	}
	curr := t.root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.val
}

func(n *BTreeNode) Write(w io.Writer) {
	w.Write([]byte(fmt.Sprintf(" 🌳%d", n.val)))
}

func (t *BTreeNode) PrintInOrder(w io.Writer) {
	if t == nil {
		return 
	}
	t.left.PrintInOrder(w)
	t.Write(w)
	t.right.PrintInOrder(w)
}

func (t *BTreeNode) PrintPreOrder(w io.Writer) {
	if t == nil {
		return
	}
	t.Write(w)	
	t.left.PrintPreOrder(w)
	t.right.PrintPreOrder(w)
}

func (t *BTreeNode) PrintPostOrder(w io.Writer) {
	if t == nil {
		return
	}
	t.left.PrintPostOrder(w)
	t.right.PrintPostOrder(w)
	t.Write(w)
}

func (t *BTree) ParseFromArray(arr []int) {
	for _, e := range arr {
		t.Insert(e)
	}
}
