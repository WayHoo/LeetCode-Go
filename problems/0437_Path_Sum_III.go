package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/path-sum-iii/
 */

func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSumWithRoot(root, targetSum) + PathSum(root.Left, targetSum) + PathSum(root.Right, targetSum)
}

func pathSumWithRoot(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val == target {
		sum++
	}
	sum += pathSumWithRoot(root.Left, target-root.Val) + pathSumWithRoot(root.Right, target-root.Val)
	return sum
}
