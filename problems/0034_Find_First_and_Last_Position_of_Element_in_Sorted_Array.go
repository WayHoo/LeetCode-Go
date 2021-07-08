package problems

/*
 LeetCode: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/

func SearchRange(nums []int, target int) []int {
	N := len(nums)
	if N == 0 || target < nums[0] || target > nums[N-1] {
		return []int{-1, -1}
	}
	start := findIndex(nums, target, true)
	if start == -1 {
		return []int{-1, -1}
	}
	end := findIndex(nums, target, false)
	return []int{start, end}
}

func findIndex(nums []int, target int, findStart bool) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if findStart {
			if nums[m] >= target {
				r = m
			} else {
				l = m + 1
			}
		} else {
			if m == l {
				m++
			}
			if nums[m] <= target {
				l = m
			} else {
				r = m - 1
			}
		}
	}
	if nums[l] != target {
		l = -1
	}
	return l
}
