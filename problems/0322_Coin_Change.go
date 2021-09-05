package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/coin-change/
 */

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] == -1 {
				continue
			}
			if dp[i] <= 0 {
				dp[i] = dp[i-coin] + 1
			} else {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	return dp[amount]
}
