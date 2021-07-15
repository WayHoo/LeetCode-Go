package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/unique-paths/
 */

func NumberOfArithmeticSlices(nums []int) int {
    if len(nums) < 3 {
        return 0
    }
    cnt, length := 0, 2
    for i := 2; i < len(nums); i++ {
        if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
            length++
        } else {
            if length >= 3 {
                cnt += (length-2)*(length-1)/2
            }
            length = 2
        }
    }
    if length >= 3 {
        cnt += (length-2)*(length-1)/2
    }
    return cnt
}
