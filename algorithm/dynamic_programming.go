package algorithm

/**
 * Reference: https://zhuanlan.zhihu.com/p/93857890
 */

// KnapsackProblem01 0-1背包
/**
 * W 为背包总体积
 * N 为物品数量
 * weights 数组存储 N 个物品的重量
 * values 数组存储 N 个物品的价值
 */
func Knapsack01(W, N int, weights, values []int) int {
	dp := make([]int, W+1)
	for i := 1; i <= N; i++ {
		w, v := weights[i-1], values[i-1]
		// 滚动数组只能逆向覆盖
		for j := W; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	return dp[W]
}

// UnboundedKnapsack 完全背包问题，每种物品数量无限制
func UnboundedKnapsack(W, N int, weights, values []int) int {
	dp := make([]int, W+1)
	for i := 1; i <= N; i++ {
		w, v := weights[i-1], values[i-1]
		// 滚动数组只能正向覆盖
		for j := w; j <= W; j++ {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	return dp[W]
}

// BoundedKnapsack 多重背包问题，每种物品有限制
func BoundedKnapsack(W, N int, weights, values, counts []int) int {
	dp := make([]int, W+1)
	for i := 1; i <= N; i++ {
		w, v, n := weights[i-1], values[i-1], counts[i-1]
		// 滚动数组只能逆向覆盖
		for j := W; j >= w; j-- {
			for k := 0; k <= min(n, j/w); k++ {
				dp[j] = max(dp[j], dp[j-k*w]+k*v)
			}
		}
	}
	return dp[W]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
