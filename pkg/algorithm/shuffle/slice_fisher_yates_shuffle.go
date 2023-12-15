package shuffle

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func SliceFisherYatesShuffle[T constraints.Ordered](values []T) {
	for i := 0; i < len(values)-1; i++ {
		j := rand.Intn(i + 2)
		if i != j {
			values[i], values[j] = values[j], values[i]
		}
	}
}
