package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/linked_list/single_linked_list"
	"golang.org/x/exp/constraints"
)

func SingleLinkedListInsertionSort[T constraints.Ordered](node *single_linked_list.Node[T], ascending bool) *single_linked_list.Node[T] {
	var previous *single_linked_list.Node[T]
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
			var left *single_linked_list.Node[T]
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
					cNext := current.Next

					// disconnect current from previous
					if previous != nil {
						previous.Next = cNext
					}

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
