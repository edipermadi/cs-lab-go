package test

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func SingleLinkedListFromSlice[T constraints.Ordered](values []T) *list.SingleLink[T] {
	if len(values) == 0 {
		return nil
	}

	head := list.NewSingleLink(values[0])
	node := head
	for i := 1; i < len(values); i++ {
		node = node.Append(list.NewSingleLink(values[i]))
	}

	return head
}

func SliceFromSingleLinkedList[T constraints.Ordered](node *list.SingleLink[T]) []T {
	entries := make([]T, 0)
	for n := node; n != nil; n = n.Next {
		entries = append(entries, n.Value)
	}

	return entries
}
