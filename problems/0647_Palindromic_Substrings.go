package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/palindromic-substrings/
 */

func CountSubstrings(s string) int {
	cnt := 0
	list := []rune(s)
	for i := 0; i < len(list); i++ {
		cnt += extendSubStrings(list, i, i)
		cnt += extendSubStrings(list, i, i+1)
	}
	return cnt
}

func extendSubStrings(list []rune, l, r int) (count int) {
	for ; l >= 0 && r < len(list) && list[l] == list[r]; l, r = l-1, r+1 {
		count++
	}
	return
}
