package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/minimum-window-substring/
 */

func MinWindow(s string, t string) string {
	ans := s + " "
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}
	sMap := make(map[byte]int)
	for lIdx, rIdx := 0, 0; rIdx < len(s); rIdx++ {
		sMap[s[rIdx]]++
		for isFit(sMap, tMap) {
			if rIdx-lIdx+1 < len(ans) {
				ans = s[lIdx : rIdx+1]
			}
			sMap[s[lIdx]]--
			lIdx++
		}
	}
	if len(ans) <= len(s) {
		return ans
	}
	return ""
}

func isFit(source, target map[byte]int) bool {
	for k, v := range target {
		if source[k] < v {
			return false
		}
	}
	return true
}
