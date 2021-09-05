package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/wiggle-subsequence/
 */

// WiggleMaxLength 跳过非严格单调的数字即可
func WiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	ans := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		ans = 2
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && preDiff <= 0 || diff < 0 && preDiff >= 0 {
			ans++
			preDiff = diff
		}
	}
	return ans
}
