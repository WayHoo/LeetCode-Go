package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/subsets/
 */

// SubsetsDFS DFS解法
func SubsetsDFS(nums []int) [][]int {
	var set []int
	var ans [][]int
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		set = append(set, nums[idx])
		dfs(idx + 1)
		set = set[:len(set)-1]
		dfs(idx + 1)
	}
	dfs(0)
	return ans
}

// SubsetsBits 位运算解法
func SubsetsBits(nums []int) [][]int {
	ans := make([][]int, 0)
    for mask := 1<<len(nums) - 1; mask >= 0; mask-- {
        var tmp []int
        for i, num := range nums {
            if mask&(1<<i) > 0 {
                tmp = append(tmp, num)
            }
        }
        ans = append(ans, tmp)
	}
	return ans
}
