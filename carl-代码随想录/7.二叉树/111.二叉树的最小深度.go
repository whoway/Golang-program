/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func DFS(root *TreeNode, depth int) {
	if nil == root.Left && nil == root.Right {
		res = min(res, depth)
		return
	}

	if nil != root.Left {
		DFS(root.Left, depth+1)
	}

	if nil != root.Right {
		DFS(root.Right, depth+1)
	}
}

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	res = math.MaxInt
	DFS(root, 1)
	return res
}
