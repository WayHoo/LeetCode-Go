package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/sort-colors/
 */

// CountPrimes 素数筛法
func CountPrimes(n int) int {
    primeMap := make([]bool, n)
    cnt := 0
    for i := 2; i < n; i++ {
        if !primeMap[i] {
            cnt++
            for k := i + i; k < n; k += i {
                primeMap[k] = true
            }
        }
    }
    return cnt
}
