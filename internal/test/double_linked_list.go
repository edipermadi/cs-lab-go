package test

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func DoubleLinkedListFromSlice[T constraints.Ordered](values []T) *list.DoubleLink[T] {
	if len(values) == 0 {
		return nil
	}

	head := list.NewDoubleLink(values[0])
	node := head
	for i := 1; i < len(values); i++ {
		node = node.Append(list.NewDoubleLink(values[i]))
	}

	return head
}

func SliceFromDoubleLinkedList[T constraints.Ordered](node *list.DoubleLink[T]) []T {
	entries := make([]T, 0)
	for n := node; n != nil; n = n.Next {
		entries = append(entries, n.Value)
	}

	return entries
}
