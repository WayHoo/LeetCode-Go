package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/counting-bits/
 */

func CountBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	bits := make([]int, n+1)
	bits[0], bits[1] = 0, 1
	base := 1
	for i := 2; i <= n; i++ {
		if i == base<<1 {
			base <<= 1
			bits[i] = 1
		} else {
			bits[i] = bits[i-base] + 1
		}
	}
	return bits
}
