package problems

/*
 LeetCode: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
*/

func MaxProfit(prices []int) int {
	N := len(prices)
	if N <= 1 {
		return 0
	}
	maxPrice, maxProfit := prices[N-1], 0
	for i := N - 2; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else if prices[i] < maxPrice {
			delta := maxPrice - prices[i]
			if delta > maxProfit {
				maxProfit = delta
			}
		}
	}
	return maxProfit
}
