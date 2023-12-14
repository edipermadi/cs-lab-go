package double_linked_list

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Value    T
	Previous *Node[T]
	Next     *Node[T]
}

func (n *Node[T]) Head() *Node[T] {
	for node := n; ; {
		if node.Previous == nil {
			return node
		}

		node = node.Previous
	}
}

func (n *Node[T]) Tail() *Node[T] {
	for node := n; ; {
		if node.Next == nil {
			return node
		}

		node = node.Next
	}
}

func (n *Node[T]) ToSlice() []T {
	values := make([]T, 0)
	for node := n; ; {
		if node == nil {
			break
		}

		values = append(values, node.Value)
		node = node.Next
	}

	return values
}

func (n *Node[T]) Length() int {
	count := 0
	for node := n.Head(); ; {
		count++
		if node.Next == nil {
			break
		}

		node = node.Next
	}

	return count
}

func FromSlice[T constraints.Ordered](values []T) *Node[T] {
	var prev, head *Node[T]
	for i := 0; i < len(values); i++ {
		node := &Node[T]{Value: values[i]}
		if head == nil {
			head = node
		}

		if prev != nil {
			prev.Next = node
			node.Previous = prev
		}

		prev = node
	}

	return head
}
