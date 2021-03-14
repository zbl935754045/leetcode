package main


//前中后序遍历的万能公式
func erchashubianli(root *TreeNode) []int {
	var pre func(node *TreeNode)
	vals := make([]int, 0)
	pre = func(node *TreeNode) {
		if node != nil {
			//前序
			vals = append(vals, node.Val)
			pre(node.Left)
			pre(node.Right)

			//中序
			pre(node.Left)
			vals = append(vals, node.Val)
			pre(node.Right)

			//后序
			pre(node.Left)
			pre(node.Right)
			vals = append(vals, node.Val)
		}
	}
	pre(root)
	return vals
}
