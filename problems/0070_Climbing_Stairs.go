package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/climbing-stairs/
 */

func ClimbStairs(n int) int {
    if n <= 2 {
        return n
    }
    pre2, pre1 := 1, 2
    for i := 3; i <= n; i++ {
        cur := pre2 + pre1
        pre2, pre1 = pre1, cur
    }
    return pre1
}
