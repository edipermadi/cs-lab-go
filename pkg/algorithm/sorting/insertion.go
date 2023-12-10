package sorting

import (
	"github.com/edipermadi/cs-lab-go/pkg/data_structure/linked_list/double_linked_list"
	"golang.org/x/exp/constraints"
)

func SliceInsertionSort[T constraints.Ordered](values []T, ordering Ordering) {
	for i := 0; i < len(values); i++ {
		var left T
		current := values[i]

		// scan values left to right
		for j := 0; j <= i; j++ {
			right := values[j]

			process := false
			if ordering == Ascending {
				// find index where left < current < right
				process = (j == 0 || left < current) && current < right
			} else if ordering == Descending {
				// find index where left > current > right
				process = (j == 0 || left > current) && current > right
			}

			// when condition met, shift and insert value
			if process {
				// shift right unordered values
				for k := i - 1; k >= j; k-- {
					values[k+1] = values[k]
				}

				// insert value
				values[j] = current
			}

			// update next left
			left = right
		}
	}
}

func DoubleLinkedListInsertionSort[T constraints.Ordered](node *double_linked_list.Node[T], ordering Ordering) *double_linked_list.Node[T] {
	head := node.Head()
	for current := head; ; {
		if current == nil {
			break
		}

		// check if value is sorted
		notSorted := false
		if ordering == Ascending {
			notSorted = current.Previous != nil && current.Previous.Value > current.Value
		} else if ordering == Descending {
			notSorted = current.Previous != nil && current.Previous.Value < current.Value
		}

		if notSorted {
			var left *double_linked_list.Node[T]
			for right := head; ; {
				if right == nil {
					break
				}

				// check if value is sorted
				process := false
				if ordering == Ascending {
					process = (left == nil || left.Value < current.Value) && current.Value < right.Value
				} else if ordering == Descending {
					process = (left == nil || left.Value > current.Value) && current.Value > right.Value
				}

				if process {
					cNext := current.Next
					cPrev := current.Previous

					// disconnect current from previous
					if current.Previous != nil {
						current.Previous.Next = cNext
					}

					// disconnect current from next
					if current.Next != nil {
						current.Next.Previous = cPrev
					}

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
