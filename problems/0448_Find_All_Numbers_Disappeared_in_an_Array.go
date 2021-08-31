package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
 */

func FindDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 || nums[i] == nums[nums[i]-1] {
			i++
			continue
		}
		idx := nums[i] - 1
		nums[i], nums[idx] = nums[idx], nums[i]
	}
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
