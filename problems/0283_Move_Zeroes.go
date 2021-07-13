package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/move-zeroes/
 */

func MoveZeroes(nums []int)  {
	idx, N := 0, len(nums)
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}
	for ;idx < N; idx++ {
		nums[idx] = 0
	}
}
