package problems

/*
 LeetCode: https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
*/

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}
		m := l + (r-l)/2
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[r]
}
