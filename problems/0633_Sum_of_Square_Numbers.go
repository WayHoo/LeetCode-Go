package problems

import "math"

/**
 * LeetCode: https://leetcode-cn.com/problems/sum-of-square-numbers/
 */

func JudgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	x := int(math.Sqrt(float64(c)))
	left, right := 0, x
	for left <= right {
		val := left * left + right * right
		if val == c {
			return true
		} else if val < c {
			left++
		} else {
			right--
		}
	}
	return false
}
