package tree

type Index int

func (i Index) Level() int {
	if i < 1 {
		return 0
	}

	for j := 0; j < 32; j++ {
		val := (1 << j) - 1
		if int(i) < val {
			return j - 1
		}
	}

	return 0
}

func (i Index) Parent() Index {
	if i < 1 {
		return 0
	}

	return (i - 1) / 2
}

func (i Index) LeftChild() Index {
	if i < 0 {
		return 0
	}

	return (i * 2) + 1
}

func (i Index) RightChild() Index {
	if i < 0 {
		return 0
	}

	return (i * 2) + 2
}

func (i Index) Sibling() Index {
	switch {
	case i < 1:
		return 0
	case i&1 > 0:
		return i + 1
	default:
		return i - 1
	}
}

func (i Index) Leftmost() Index {
	if i < 1 {
		return 0
	}

	j := i.Level()
	return Index((1 << j) - 1)
}

func (i Index) Rightmost() Index {
	if i < 1 {
		return 0
	}

	j := i.Level() + 1
	return Index((1 << j) - 2)
}

func (i Index) Children() []Index {
	return []Index{i.LeftChild(), i.RightChild()}
}
