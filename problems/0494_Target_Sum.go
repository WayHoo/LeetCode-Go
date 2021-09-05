package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/target-sum/
 */

func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || sum < -target || (sum+target)%2 == 1 {
		return 0
	}
	W := (sum + target) / 2
	dp := make([]int, W+1)
	dp[0] = 1
	for _, num := range nums {
		for j := W; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[W]
}

func FindTargetSumWaysDFS(nums []int, target int) int {
	val, cnt := 0, 0
	var dfs func(int)
	dfs = func(idx int) {
		if idx >= len(nums) {
			if val == target {
				cnt++
			}
			return
		}
		val += nums[idx]
		dfs(idx + 1)
		val -= 2 * nums[idx]
		dfs(idx + 1)
		val += nums[idx]
	}
	dfs(0)
	return cnt
}
