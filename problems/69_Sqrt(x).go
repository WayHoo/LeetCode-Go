package problems

/*
 LeetCode: https://leetcode-cn.com/problems/sqrtx/description/
*/

func MySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	l, r := 1, x
	for l <= r {
		m := l + (r - l) / 2
		sqrt := x / m
		if sqrt == m {
			return m
		} else if sqrt < m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}
