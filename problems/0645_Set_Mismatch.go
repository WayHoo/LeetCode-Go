package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/set-mismatch/
 */

func FindErrorNums(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return []int{nums[i], i+1}
		}
	}
	return nil
}