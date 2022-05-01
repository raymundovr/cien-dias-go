package datastructs

import "fmt"

type Valuable interface {
	int | float64 | string
}

type BinaryTree[T Valuable] struct {
	value		T
	left,right	*BinaryTree[T]
}

func NewBinaryTree[T Valuable](value T) *BinaryTree[T] {
	return &BinaryTree[T]{ value: value, left: nil, right: nil }
}

func Insert[T Valuable](tree *BinaryTree[T], value T) *BinaryTree[T] {
	if tree == nil {
		tree = new(BinaryTree[T])
		tree.value = value
		return tree
	}

	if value < tree.value {
		tree.left = Insert(tree.left, value)
	} else {
		tree.right = Insert(tree.right, value)
	}

	return tree
}

func Inorder[T Valuable](tree *BinaryTree[T]) {
	if tree == nil {
		return ;
	}
	Inorder(tree.left)
	fmt.Println("Value ", tree.value)
	Inorder(tree.right)
}

