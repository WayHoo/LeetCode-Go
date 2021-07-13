package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
 */

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for row, col := 0, n - 1; row < m && col >= 0; {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
