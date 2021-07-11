package problems

/*
 LeetCode: https://leetcode-cn.com/problems/is-subsequence/
*/

func CheckPossibility(nums []int) bool {
	changed := false
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			if changed {
				return false
			}
			if i - 1 < 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}
			changed = true
		}
	}
	return true
}
