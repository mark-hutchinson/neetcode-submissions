/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	return traverse(root, 0, result)
}

func traverse(root *TreeNode, level int, result [][]int) [][]int{
	if root == nil {
		return result
	}
	if len(result) == level {
		// add new entry for this level's results
		result = append(result, make([]int, 0))
	}

	result[level] = append(result[level], root.Val)

	// traverse next level
	result = traverse(root.Left, level+1, result)
	result = traverse(root.Right, level+1, result)
	return result
}
