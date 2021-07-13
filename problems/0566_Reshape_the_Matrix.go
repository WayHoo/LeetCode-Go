package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/reshape-the-matrix/
 */

// MatrixReshape 速度更快
func MatrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	x, y := 0, 0
	ret := make([][]int, r)
	for i := 0; i < r; i++ {
		ret[i] = make([]int, c)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if y == c {
				x, y = x + 1, 0
			}
			ret[x][y] = mat[i][j]
			y++
		}
	}
	return ret
}

// MatrixReshapeV2 代码简洁，但包含除法和求余计算，速度稍慢
func MatrixReshapeV2(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ret := make([][]int, r)
	for i := 0; i < r; i++ {
		ret[i] = make([]int, c)
	}
	idx := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ret[i][j] = mat[idx/n][idx%n]
			idx++
		}
	}
	return ret
}