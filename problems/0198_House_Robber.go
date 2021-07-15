package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/house-robber/
 */

func Rob(nums []int) int {
    pre2, pre1 := 0, nums[0]
    for i := 1; i < len(nums); i++ {
        cur := nums[i] + pre2
        if cur < pre1 {
            cur = pre1
        }
        pre2, pre1 = pre1, cur
    }
    return pre1
}
