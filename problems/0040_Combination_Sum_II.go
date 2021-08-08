package problems

import "sort"

/**
 * LeetCode: https://leetcode-cn.com/problems/combination-sum-ii/
 */

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	visited := make([]bool, len(candidates))
	var combination func(start, target int) [][]int
	combination = func(start, target int) [][]int {
		ret := make([][]int, 0)
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if candidates[i] == target {
				ret = append(ret, []int{candidates[i]})
				break
			}
			if i > 0 && candidates[i] == candidates[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			if ans := combination(i+1, target-candidates[i]); ans != nil {
				for _, tmp := range ans {
					ret = append(ret, append(tmp, candidates[i]))
				}
			}
			visited[i] = false
		}
		return ret
	}
	return combination(0, target)
}
