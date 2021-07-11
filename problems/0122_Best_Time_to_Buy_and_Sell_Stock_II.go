package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
 */

func MaxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		if delta > 0 {
			profit += delta
		}
	}
	return profit
}
