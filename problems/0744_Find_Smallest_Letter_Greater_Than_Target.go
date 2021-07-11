package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/sqrtx/
 */

func NextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters) - 1
	for l <= r {
		m := l + (r - l) / 2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l < len(letters) {
		return letters[l]
	}
	return letters[0]
}
