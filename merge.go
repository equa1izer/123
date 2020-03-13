type TreeNodeM struct {
	Left  *TreeNodeM
	Val   int
	Right *TreeNodeM
}

// MergeTrees q
func MergeTrees(t1 *TreeNodeM, t2 *TreeNodeM) *TreeNodeM {

	newRoot := &TreeNodeM{}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	newRoot.Val = t1.Val + t2.Val
	newRoot.Left = MergeTrees(t1.Left, t2.Left)
	newRoot.Right = MergeTrees(t1.Right, t2.Right)

	return newRoot
}

// PrintAnotherTree q
func PrintAnotherTree(root *TreeNodeM) {
	if root != nil {
		PrintAnotherTree(root.Left)
		fmt.Println(root.Val)
		PrintAnotherTree(root.Right)
	}
}

// InsertAnotherTNode q
func InsertAnotherTNode(root *TreeNodeM, data int) *TreeNodeM {
	newNode := &TreeNodeM{Val: data}

	if root == nil {
		return newNode
	}

	if data < root.Val {
		// fmt.Println(root.Left)
		root.Left = InsertAnotherTNode(root.Left, data)
	} else {
		// fmt.Println(root.Right)
		root.Right = InsertAnotherTNode(root.Right, data)
	}

	return root
}
