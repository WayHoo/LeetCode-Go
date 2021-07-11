package problems

/*
 LeetCode: https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
*/

func ReverseVowels(s string) string {
	if len(s) == 0 {
		return ""
	}
	charMap := map[int32]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	r := []rune(s)
	left, right := 0, len(s)-1
	for left < right {
		if charMap[r[left]] && charMap[r[right]] {
			tmp := r[left]
			r[left] = r[right]
			r[right] = tmp
			left++
			right--
		}
		if !charMap[r[left]] {
			left++
		}
		if !charMap[r[right]] {
			right--
		}
	}
	return string(r)
}
