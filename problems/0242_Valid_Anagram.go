package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/valid-anagram/
 */

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[int32]uint16)
	for _, c := range s {
		charMap[c]++
	}
	for _, c := range t {
		if _, ok := charMap[c]; ok {
			if charMap[c] > 0 {
				charMap[c]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}