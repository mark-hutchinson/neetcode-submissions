/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, balanced := depth(root)
	return balanced
}

func depth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	var balanced bool
	leftDepth, balanced := depth(root.Left)
	if !balanced {
		return 0, false // don't care about height
	}
	rightDepth, balanced := depth(root.Right)
	if !balanced {
		return 0, false
	}

	height := max(leftDepth, rightDepth) + 1

	delta := leftDepth - rightDepth
	if delta < 0 {
		delta = delta * -1
	}

	return height, delta <= 1
}