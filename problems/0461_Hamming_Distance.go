package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/hamming-distance/
 */

func HammingDistance(x int, y int) int {
	ans := 0
	z := x ^ y
	for z != 0 {
		if z%2 == 1 {
			ans++
		}
		z >>= 1
	}
	return ans
}
