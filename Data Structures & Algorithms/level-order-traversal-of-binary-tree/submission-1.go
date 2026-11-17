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
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue) 
		levelVals := make([]int, levelSize)
		for i:= 0; i<levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelVals[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelVals)
	}

	return result
}

// func traverse(root *TreeNode, level int, result [][]int) [][]int{
// 	if root == nil {
// 		return result
// 	}
// 	if len(result) == level {
// 		// add new entry for this level's results
// 		result = append(result, make([]int, 0))
// 	}

// 	result[level] = append(result[level], root.Val)

// 	// traverse next level
// 	result = traverse(root.Left, level+1, result)
// 	result = traverse(root.Right, level+1, result)
// 	return result
// }
