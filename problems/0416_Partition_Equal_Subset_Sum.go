package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/partition-equal-subset-sum/
 */

func CanPartition(nums []int) bool {
	sum, maxVal := 0, 0
	for _, num := range nums {
		sum += num
		if num > maxVal {
			maxVal = num
		}
	}
	if sum%2 == 1 || maxVal > sum/2 {
		return false
	}
	target := sum / 2
	if maxVal == target {
		return true
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
