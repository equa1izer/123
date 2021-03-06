func Reverse(node *NodeAddL) *NodeAddL {
	var copy *NodeAddL
	for it := node; it != nil; it = it.Next {
		copy = PushFront(copy, it.Num)
	}
	return copy
}

// PushFront q
func PushFront(node *NodeAddL, data int) *NodeAddL {
	return &NodeAddL{Num: data, Next: node}
}
