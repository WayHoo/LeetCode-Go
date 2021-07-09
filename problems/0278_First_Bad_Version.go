package problems

/*
 LeetCode: https://leetcode-cn.com/problems/first-bad-version/description/
*/

func FirstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

func isBadVersion(version int) bool {
	return false
}