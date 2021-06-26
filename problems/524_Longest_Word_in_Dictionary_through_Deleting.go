package problems

import (
	"sort"
	"strings"
)

/*
 LeetCode: https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/description/
*/

func FindLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		} else {
			return strings.Compare(dictionary[i], dictionary[j]) <= 0
		}
	})
	for _, str := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(str) {
			if s[i] == str[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(str) {
			return str
		}
	}
	return ""
}
