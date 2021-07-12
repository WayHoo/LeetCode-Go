package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/isomorphic-strings/
 */

// IsIsomorphic 哈希解法
func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	runeMap := make(map[rune]rune)
	mappedRune := make(map[rune]bool)
	rs, rt := []rune(s), []rune(t)
	for i := 0; i < len(rs); i++ {
		if r, ok := runeMap[rs[i]]; ok {
			if r != rt[i] {
				return false
			}
		} else if !mappedRune[rt[i]] {
			runeMap[rs[i]] = rt[i]
			mappedRune[rt[i]] = true
		} else {
			return false
		}
	}
	return true
}

// IsIsomorphicV2 记录一个字符上次出现的位置，如果两个字符串中的字符上次出现的位置一样，那么就属于同构。
func IsIsomorphicV2(s string, t string) bool {
	preIdxS, preIdxT := make(map[rune]int), make(map[rune]int)
	rs, rt := []rune(s), []rune(t)
	for i := 0; i < len(rs); i++ {
		if preIdxS[rs[i]] == preIdxT[rt[i]] {
			preIdxS[rs[i]] = i + 1
			preIdxT[rt[i]] = i + 1
		} else {
			return false
		}
	}
	return true
}
