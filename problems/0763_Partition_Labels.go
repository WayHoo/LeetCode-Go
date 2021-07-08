package problems

/*
 LeetCode: https://leetcode-cn.com/problems/partition-labels/description/
*/

func PartitionLabels(s string) []int {
	charIndex := make([]int, 26)
	for i, ch := range s {
		charIndex[ch-'a'] = i
	}
	res := make([]int, 0)
	for i := 0; i < len(s); {
		endIdx := charIndex[s[i]-'a']
		for j := i + 1; j <= endIdx; j++ {
			if charIndex[s[j]-'a'] > endIdx {
				endIdx = charIndex[s[j]-'a']
			}
		}
		res = append(res, endIdx-i+1)
		i = endIdx + 1
	}
	return res
}
