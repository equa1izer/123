type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

// InsertTNode q
func InsertTNode(root *TNode, data int) *TNode {
	newNode := &TNode{Val: data}

	if root == nil {
		return newNode
	}

	if data < root.Val {
		root.Left = InsertTNode(root.Left, data)
	} else {
		root.Right = InsertTNode(root.Right, data)
	}

	return root
}

// PrintTree q
func PrintTree(root *TNode) {
	if root != nil {
		PrintTree(root.Left)
		fmt.Println(root.Val)
		PrintTree(root.Right)
	}
}

// InvertTree q
func InvertTree(root *TNode) *TNode {

	if root == nil {
		return root
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}

// Print2DTree q
func Print2DTree(root *TNode, space int) {
	const count = 6

	// Base case
	if root == nil {
		return
	}

	// Increase distance between levels
	space += count

	// Process right child first
	Print2DTree(root.Right, space)

	// Print current node after space
	// count
	fmt.Print("\n")
	for i := count; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Print(root.Val, "\n")

	// Process left child
	Print2DTree(root.Left, space)
}
