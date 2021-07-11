package problems

import (
	"math"
)

/*
 LeetCode: https://leetcode-cn.com/problems/perfect-squares/
*/

func NumSquaresDP(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		for j := 1; j * j <= i;j++ {
			if dp[i-j*j] < min {
				min = dp[i-j*j]
			}
			dp[i] = min + 1
		}
	}
	return dp[n]
}

func NumSquaresBFS(n int) int {
	squares := genSquares(n)
	queue := []int{n}
	visited := make([]bool, n+1)
	level, curLevelNum, nextLevelNum := 1, 1, 0
	for len(queue) > 0 {
		cur := queue[0]
		for _, val := range squares {
			if delta := cur - val; delta >= 0 && !visited[delta] {
				if delta == 0 {
					return level
				}
				queue = append(queue, delta)
				visited[delta] = true
				nextLevelNum++
			}
		}
		if curLevelNum--; curLevelNum == 0 {
			level++
			curLevelNum, nextLevelNum = nextLevelNum, 0
		}
		queue = queue[1:]
	}
	return 0
}

func genSquares(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		val := i * i
		if val <= n {
			res = append(res, val)
		} else {
			break
		}
	}
	return res
}
