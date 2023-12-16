package list

import "golang.org/x/exp/constraints"

func NewDoubleLink[T constraints.Ordered](value T) *DoubleLink[T] {
	return &DoubleLink[T]{Value: value}
}

type DoubleLink[T constraints.Ordered] struct {
	Value    T
	Previous *DoubleLink[T]
	Next     *DoubleLink[T]
}

func (d *DoubleLink[T]) Head() *DoubleLink[T] {
	for node := d; ; {
		if node.Previous == nil {
			return node
		}

		node = node.Previous
	}
}

func (d *DoubleLink[T]) Tail() *DoubleLink[T] {
	for node := d; ; {
		if node.Next == nil {
			return node
		}

		node = node.Next
	}
}

func (d *DoubleLink[T]) Length() int {
	count := 0
	for node := d.Head(); ; {
		count++
		if node.Next == nil {
			break
		}

		node = node.Next
	}

	return count
}

func (d *DoubleLink[T]) Prepend(x *DoubleLink[T]) *DoubleLink[T] {
	x.Next = d
	d.Previous = x
	return x
}

func (d *DoubleLink[T]) Append(x *DoubleLink[T]) *DoubleLink[T] {
	d.Next = x
	x.Previous = d
	return x
}

func (d *DoubleLink[T]) Splice() {
	previous, next := d.Previous, d.Next

	if d.Previous != nil {
		d.Previous.Next = next
	}

	if d.Next != nil {
		d.Next.Previous = previous
	}

	d.Previous, d.Next = nil, nil
}
