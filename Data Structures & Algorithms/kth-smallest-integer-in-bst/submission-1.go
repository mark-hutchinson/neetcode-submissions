/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var results []int
	results = dfs(root, k, results)
	return results[k - 1]
}

func dfs(root *TreeNode, k int, results []int) []int{
	if root == nil {
		return results
	}
	if len(results) >= k {
		return results
	}
	// visit left
	results = dfs(root.Left, k, results)

	// visit this node
	results = append(results, root.Val)

	// visit right
	results = dfs(root.Right, k, results)
	return results
}

