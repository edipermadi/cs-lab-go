package tree

import "golang.org/x/exp/constraints"

func NewNode[T constraints.Ordered]() *Node[T] {
	return &Node[T]{}
}

type Node[T constraints.Ordered] struct {
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

func (n *Node[T]) AddLeft(left *Node[T]) *Node[T] {
	n.Left = left
	left.Parent = n
	return left
}

func (n *Node[T]) AddRight(right *Node[T]) *Node[T] {
	n.Right = right
	right.Parent = n
	return right
}
