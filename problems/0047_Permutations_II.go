package problems

import "sort"

/*
 LeetCode: https://leetcode-cn.com/problems/permutations-ii/
*/

func PermuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		var subNums []int
		subNums = append(subNums, nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)
		list := PermuteUnique(subNums)
		for idx := range list {
			res = append(res, append(list[idx], nums[i]))
		}
	}
	return res
}
