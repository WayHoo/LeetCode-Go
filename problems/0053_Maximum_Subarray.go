package problems

/*
 LeetCode: https://leetcode-cn.com/problems/maximum-subarray/
*/

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, preSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if preSum > 0 {
			preSum += nums[i]
		} else {
			preSum = nums[i]
		}
		if preSum > maxSum {
			maxSum = preSum
		}
	}
	return maxSum
}
