package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/combination-sum/
 */

func CombinationSum(candidates []int, target int) [][]int {
	var backtrack func(idx, val int) [][]int
	backtrack = func(idx, val int) [][]int {
		if idx >= len(candidates) || val < 0 {
			return nil
		}
		ret := make([][]int, 0)
		for i := idx; i < len(candidates); i++ {
			num := candidates[i]
			if num == val {
				ret = append(ret, []int{num})
				continue
			}
			if tmp := backtrack(i, val-num); len(tmp) > 0 {
				for j := range tmp {
					ret = append(ret, append([]int{num}, tmp[j]...))
				}
			}
		}
		return ret
	}
	return backtrack(0, target)
}
