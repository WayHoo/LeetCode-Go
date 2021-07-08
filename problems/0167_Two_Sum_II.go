package problems

/*
 LeetCode: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
 */

func TwoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return nil
	}
	var res []int
	left, right := 0, len(numbers)-1
	for left < right {
		val := numbers[left] + numbers[right]
		if val < target {
			left++
		} else if val > target {
			right--
		} else {
			res = []int{left+1, right+1}
			break
		}
	}
	return res
}
