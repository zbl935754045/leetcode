//给定一个二叉树的根节点 root ，返回它的 中序 遍历。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[2,1]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 哈希表 
// 👍 877 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//第一种解法
var vals94 = make([]int, 0)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return vals94
	}

	if root.Left != nil {
		inorderTraversal(root.Left)
	}
	vals94 = append(vals94, root.Val)
	if root.Right != nil {
		inorderTraversal(root.Right)
	}
	return vals94
}


//第二种写法
func inorderTraversal1(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	getInorgerVals(root, &vals)
	return vals
}

func getInorgerVals(root *TreeNode, vals *[]int) {
	if root != nil {
		getInorgerVals(root.Left, vals)
		*vals = append(*vals, root.Val)
		getInorgerVals(root.Right, vals)
	}
}

func inorderTraversal2(root *TreeNode) []int {
	var preOrder func(*TreeNode)
	vals := make([]int, 0)
	preOrder = func(node *TreeNode) {
		if node != nil {
			preOrder(node.Left)
			vals = append(vals, node.Val)
			preOrder(node.Right)
		}
	}
	preOrder(root)
	return nil
}



//leetcode submit region end(Prohibit modification and deletion)
func main(){

}