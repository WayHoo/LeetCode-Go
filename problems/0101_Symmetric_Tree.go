package problems

import "math"

/**
 * LeetCode: https://leetcode-cn.com/problems/symmetric-tree/
 */

// IsSymmetric 递归方式
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var recursion func(left, right *TreeNode) bool
	recursion = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left != nil && right != nil && left.Val == right.Val {
			return recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
		}
		return false
	}
	return recursion(root.Left, root.Right)
}

// IsSymmetricIter 迭代方式
func IsSymmetricIter(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevelNum := len(queue)
		nums := make([]int, 0)
		for i := 0; i < curLevelNum; i++ {
			if queue[i].Left == nil {
				nums = append(nums, math.MaxInt32)
			} else {
				nums = append(nums, queue[i].Left.Val)
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right == nil {
				nums = append(nums, math.MaxInt32)
			} else {
				nums = append(nums, queue[i].Right.Val)
				queue = append(queue, queue[i].Right)
			}
		}
		for i := 0; i < len(nums)/2; i++ {
			if nums[i] != nums[len(nums)-1-i] {
				return false
			}
		}
		queue = queue[curLevelNum:]
	}
	return true
}
