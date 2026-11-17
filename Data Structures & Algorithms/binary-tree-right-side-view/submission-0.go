/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)
	levelSize := 1

	var result []int
	for len(queue) > 0 {
		result = append(result, queue[0].Val)

		for i := 0; i < levelSize; i++ {
			// pop off from queue
			node := queue[0]
			queue = queue[1:]
			// push right
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		levelSize = len(queue)

	}

	return result
}
