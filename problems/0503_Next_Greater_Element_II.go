package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/next-greater-element-ii/
 */

// NextGreaterElements 把循环数组「拉直」，即复制该序列的前 n-1n−1 个元素拼接在原序列的后面。实现上采用取模操作。
func NextGreaterElements(nums []int) []int {
	N := len(nums)
	stack := make([]int, 0)
	ans := make([]int, N)
	for i := 0; i < 2*N-1; i++ {
		for len(stack) > 0 && nums[i%N] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i%N]
			stack = stack[:len(stack)-1]
		}
		if i < N {
			stack = append(stack, i)
		}
	}
	for _, idx := range stack {
		ans[idx] = -1
	}
	return ans
}
