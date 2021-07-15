package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/house-robber-ii/
 * 解题思路：根据偷不偷第一家，把问题拆分为两个无环的子问题
 */

func RobV2(nums []int) int {
	robNonCycle := func(start, end int) int {
		if start > end {
			return 0
		}
		pre2, pre1 := 0, nums[start]
		for i := start + 1; i <= end; i++ {
			cur := pre2 + nums[i]
			if cur < pre1 {
				cur = pre1
			}
			pre2, pre1 = pre1, cur
		}
		return pre1
	}
	robFirst := nums[0] + robNonCycle(2, len(nums)-2)
	notRobFirst := robNonCycle(1, len(nums)-1)
	if robFirst >= notRobFirst {
		return robFirst
	} else {
		return notRobFirst
	}
}
