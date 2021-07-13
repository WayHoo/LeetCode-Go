package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/max-consecutive-ones/
 */

func FindMaxConsecutiveOnes(nums []int) int {
	cnt, maxLen := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if cnt++; cnt > maxLen {
				maxLen = cnt
			}
		} else {
			cnt = 0
		}
	}
	return maxLen
}
