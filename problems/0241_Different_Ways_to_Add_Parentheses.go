package problems

import (
	"strconv"
)

/**
 * LeetCode: https://leetcode-cn.com/problems/different-ways-to-add-parentheses/
 */

func DiffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		if ch == '+' || ch == '-' || ch == '*' {
			left := DiffWaysToCompute(expression[:i])
			right := DiffWaysToCompute(expression[i+1:])
			for _, m := range left {
				for _, n := range right {
					switch ch {
					case '+':
						res = append(res, m+n)
					case '-':
						res = append(res, m-n)
					case '*':
						res = append(res, m*n)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		val, _ := strconv.ParseInt(expression, 10, 64)
		res = append(res, int(val))
	}
	return res
}
