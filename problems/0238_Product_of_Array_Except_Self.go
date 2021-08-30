package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/product-of-array-except-self/
 */

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	product := 1
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		product *= nums[i-1]
		ans[i] = product
	}
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= product
		product *= nums[i]
	}
	return ans
}
