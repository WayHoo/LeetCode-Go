package problems

import "math"

/**
 * LeetCode: https://leetcode-cn.com/problems/two-sum/
 */

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	var maxGain func(*TreeNode) int
	maxGain = func(tree *TreeNode) int {
		if tree == nil {
			return 0
		}
		leftGain := max(maxGain(tree.Left), 0)
		rightGain := max(maxGain(tree.Right), 0)
		sum := tree.Val + leftGain + rightGain
		maxSum = max(maxSum, sum)
		return tree.Val + max(leftGain, rightGain)
	}
	_ = maxGain(root)
	return maxSum
}
