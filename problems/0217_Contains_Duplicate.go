package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/contains-duplicate/
 */

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
