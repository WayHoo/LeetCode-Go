package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/daily-temperatures/
 */

func DailyTemperatures(temperatures []int) []int {
	stack, day := make([]int, 0), make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			day[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return day
}
