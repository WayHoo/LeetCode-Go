package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/jump-game/
 */

func CanJump(nums []int) bool {
    maxPos := 0
    for i := 0; i < len(nums) && i <= maxPos; i++ {
        if pos := i+nums[i]; pos > maxPos {
            maxPos = pos
        }
    }
    return maxPos >= len(nums)-1
}
