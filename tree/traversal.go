package tree

//指针接收者,不加*为值接收者
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}
