package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/combinations/
 */

func Combine(n int, k int) [][]int {
	return backtrack(1, n, k)
}

func backtrack(m, n, k int) [][]int {
	ret := make([][]int, 0)
	if k == 1 {
		for i := m; i <= n; i++ {
			ret = append(ret, []int{i})
		}
		return ret
	}
	for i := m; i <= n-k+1; i++ {
		tmp := backtrack(i+1, n, k-1)
		for j := range tmp {
			tmp[j] = append(tmp[j], i)
		}
		ret = append(ret, tmp...)
	}
	return ret
}
