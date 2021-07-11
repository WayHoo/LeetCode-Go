package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/two-sum/
 */

func LongestConsecutive(nums []int) int {
	m, max := make(map[int]bool), 0
	for _, num := range nums {
		m[num] = true
	}
	for num := range m {
		if m[num-1] {
			continue
		}
		cur := 1
		for i := num+1; m[i]; i, cur = i+1, cur+1 {}
		if cur > max {
			max = cur
		}
	}
	return max
}
