package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/maximum-product-subarray/
 */

func MaxProduct(nums []int) int {
	preMax, preMin, maxAns := nums[0], nums[0], nums[0]
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	for i := 1; i < len(nums); i++ {
		curMax := max(preMax*nums[i], max(preMin*nums[i], nums[i]))
		curMin := min(preMax*nums[i], min(preMin*nums[i], nums[i]))
		preMax, preMin = curMax, curMin
		if preMax > maxAns {
			maxAns = preMax
		}
	}
	return maxAns
}
