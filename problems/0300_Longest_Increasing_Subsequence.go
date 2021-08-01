package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/longest-increasing-subsequence/
 */

// LengthOfLIS 常规动态规划方法，时间复杂度 O(n^2)
func LengthOfLIS(nums []int) int {
	N := len(nums)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && (dp[j]+1) > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	max := 1
	for i := 0; i < N; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// LengthOfLISV2 改进动态规划方法，时间复杂度 O(N*logN)
func LengthOfLISV2(nums []int) int {
	N := len(nums)
	tails, idx := make([]int, N), 0
	tails[0] = nums[0]
	for i := 1; i < N; i++ {
		if nums[i] > tails[idx] {
			idx++
			tails[idx] = nums[i]
			continue
		}
		l, r := 0, idx
		for l < r {
			m := l + (r-l)/2
			if tails[m] > nums[i] {
				r = m
			} else if tails[m] < nums[i] {
				l = m + 1
			} else {
				r = m
				break
			}
		}
		tails[r] = nums[i]
	}
	return idx + 1
}
