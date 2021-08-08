package problems

import "math"

/**
 * LeetCode: https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
 */

func MinDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt32
	if root.Left != nil {
		leftDepth := MinDepthDFS(root.Left)
		if leftDepth < depth {
			depth = leftDepth
		}
	}
	if root.Right != nil {
		rightDepth := MinDepthDFS(root.Right)
		if rightDepth < depth {
			depth = rightDepth
		}
	}
	return depth + 1
}

func MinDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevelNum := len(queue)
		for i := 0; i < curLevelNum; i++ {
			cur := queue[i]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
		queue = queue[curLevelNum:]
	}
	return depth
}
