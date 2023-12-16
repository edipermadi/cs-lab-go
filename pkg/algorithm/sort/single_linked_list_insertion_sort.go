package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func SingleLinkedListInsertionSort[T constraints.Ordered](node *list.SingleLink[T], ascending bool) *list.SingleLink[T] {
	var previous *list.SingleLink[T]
	head := node
	for current := head; ; {
		if current == nil {
			break
		}

		// check if value is sorted
		notSorted := false
		if ascending {
			notSorted = previous != nil && previous.Value > current.Value
		} else {
			notSorted = previous != nil && previous.Value < current.Value
		}

		if notSorted {
			var left *list.SingleLink[T]
			for right := head; ; {
				if right == nil {
					break
				}

				// check if value is sorted
				process := false
				if ascending {
					process = (left == nil || left.Value < current.Value) && current.Value < right.Value
				} else {
					process = (left == nil || left.Value > current.Value) && current.Value > right.Value
				}

				if process {
					current.Splice(previous)

					// update current node
					current.Next = right

					// update left node
					if left != nil {
						left.Next = current
					} else {
						head = current
					}

					break
				}

				left = right
				right = right.Next
			}
		}

		previous = current
		current = current.Next
	}

	return head
}
