/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 坑：Go里面没有引用传递，Go语言是值传递！
func dfs(root *TreeNode, res *[]int) {
	if nil == root {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	dfs(root, &res)
	return res
}
