package problems

import "sort"

/**
 * LeetCode: https://leetcode-cn.com/problems/maximum-length-of-pair-chain/
 */

func FindLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] <= pairs[j][0]
	})
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(pairs)-1]
}
