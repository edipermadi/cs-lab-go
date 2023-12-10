package double_linked_list

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](node *Node[T], ascending bool) *Node[T] {
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
			var left *Node[T]
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
