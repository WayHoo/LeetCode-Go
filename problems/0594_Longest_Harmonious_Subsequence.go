package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/longest-harmonious-subsequence/
 */

func FindLHS(nums []int) int {
	m, max := make(map[int]int), 0
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if val, ok := m[k+1]; ok && v+val > max {
			max = v + val
		}
	}
	return max
}
