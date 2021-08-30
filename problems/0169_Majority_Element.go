package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/majority-element/
 */

func MajorityElement(nums []int) int {
	candi, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candi = num
		}
		if num == candi {
			count++
		} else {
			count--
		}
	}
	return candi
}
