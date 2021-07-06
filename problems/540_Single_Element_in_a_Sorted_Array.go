package problems

/*
 LeetCode: https://leetcode-cn.com/problems/single-element-in-a-sorted-array/description/
*/

func SingleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == nums[m-1] {
			if m%2 == 0 {
				r = m - 2
			} else {
				l = m + 1
			}
		} else if nums[m] == nums[m+1] {
			if m%2 == 0 {
				l = m + 2
			} else {
				r = m - 1
			}
		} else {
			return nums[m]
		}
	}
	return nums[l]
}
