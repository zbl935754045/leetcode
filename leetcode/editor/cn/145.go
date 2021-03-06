//给定一个二叉树，返回它的 后序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 
// 👍 536 👎 0

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

//简约了写法，去掉了递归前对左右子树的判空
var vals = make([]int, 0)
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return vals
	}
	postorderTraversal(root.Left)
	vals = append(vals, root.Val)
	postorderTraversal(root.Right)
	return vals
}

//leetcode submit region end(Prohibit modification and deletion)

func main(){

}