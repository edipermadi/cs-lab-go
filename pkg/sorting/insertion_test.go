package sorting_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/sorting"
	"github.com/stretchr/testify/assert"
)

func sorted(values []int) []int {
	result := append([]int{}, values...)
	sort.Ints(result)
	return result
}

func TestInsertionSort(t *testing.T) {
	values := rand.Perm(100)
	expected := sorted(values)

	sorting.InsertionSort(values)
	assert.Equal(t, expected, values)
}
