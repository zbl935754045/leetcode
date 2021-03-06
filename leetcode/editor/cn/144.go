//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
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
//输出：[1,2]
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
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 525 👎 0

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
//preorderTraversal 第一种解法
func preorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	getVals(root, &vals)
	return vals
}
func getVals(root *TreeNode, vals *[]int) {
	if root != nil {
		*vals = append(*vals, root.Val)
		getVals(root.Left, vals)
		getVals(root.Right, vals)
	}
}


//preorderTraversal1 第二种解法，不另起函数实现递归
var vals144 = make([]int, 0)
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return vals144
	}
	vals144 = append(vals144, root.Val)
	if root.Left != nil {
		preorderTraversal(root.Left)
	}
	if root.Right != nil {
		preorderTraversal(root.Right)
	}
	return vals144
}


//preorderTraversal2 第三种解法，变量函数实现
func preorderTraversal2(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}




//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
