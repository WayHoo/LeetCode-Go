package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/permutations/
 */

func Permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var res [][]int
	for i := 0; i < len(nums); i++ {
		var subNums []int
		subNums = append(subNums, nums[0:i]...)
		subNums = append(subNums, nums[i+1:]...)
		list := Permute(subNums)
		for j := range list {
			res = append(res, append(list[j], nums[i]))
		}
	}
	return res
}
