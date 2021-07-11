package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/diameter-of-binary-tree/
 */

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var depthOfBTree func(root *TreeNode) int
	depthOfBTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depthOfBTree(root.Left)
		r := depthOfBTree(root.Right)
		if l + r > maxDiameter {
			maxDiameter = l + r
		}
		if l >= r {
			return l + 1
		} else {
			return r + 1
		}
	}
	_ = depthOfBTree(root)
	return maxDiameter
}
