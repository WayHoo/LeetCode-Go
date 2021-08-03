package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/valid-parentheses/
 */

func IsValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		if s[i] == ')' && top != '(' || s[i] == '}' && top != '{' || s[i] == ']' && top != '[' {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
