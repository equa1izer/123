type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

// Changeorder q
func Changeorder(node *NodeAddL) *NodeAddL {
	var copy *NodeAddL

	it := node
	for it != nil {
		if it.Num%2 == 1 {
			copy = PushBack(copy, it.Num)
		}
		it = it.Next
	}

	it = node
	for it != nil {
		if it.Num%2 == 0 {
			copy = PushBack(copy, it.Num)
		}
		it = it.Next
	}

	return copy
}

// PushBack q
func PushBack(node *NodeAddL, x int) *NodeAddL {
	n := &NodeAddL{Num: x}

	if node == nil {
		return n
	}

	it := node

	for it.Next != nil {
		it = it.Next
	}

	it.Next = n

	return node
}
