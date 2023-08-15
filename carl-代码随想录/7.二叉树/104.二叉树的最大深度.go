/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func DFS(root *TreeNode, depth int) {
	if nil == root {
		res = max(res, depth)
		return
	}

	DFS(root.Left, depth+1)
	DFS(root.Right, depth+1)
}

func maxDepth(root *TreeNode) int {
	res = 0
	DFS(root, 0)
	return res
}
