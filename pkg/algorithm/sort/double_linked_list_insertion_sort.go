package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func DoubleLinkedListInsertionSort[T constraints.Ordered](node *list.DoubleLink[T], ascending bool) *list.DoubleLink[T] {
	head := node.Head()
	for current := head; ; {
		if current == nil {
			break
		}

		// check if value is sorted
		notSorted := false
		if ascending {
			notSorted = current.Previous != nil && current.Previous.Value > current.Value
		} else {
			notSorted = current.Previous != nil && current.Previous.Value < current.Value
		}

		if notSorted {
			var left *list.DoubleLink[T]
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
					current.Splice()

					// update current node
					current.Next = right
					current.Previous = left

					// update right node
					right.Previous = current

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

		current = current.Next
	}

	return head
}
