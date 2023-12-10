package sorting

func InsertionSort[T int | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64 | string](values []T, ordering Ordering) {
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
