package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/single-number/
 */

func SingleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
