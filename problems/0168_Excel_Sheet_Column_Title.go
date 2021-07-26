package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/excel-sheet-column-title/
 */

func ConvertToTitle(columnNumber int) string {
	var bytes []byte
	for columnNumber != 0 {
	    columnNumber--
		n := columnNumber % 26
		bytes = append(bytes, byte('A'+n))
		columnNumber /= 26
	}
	size := len(bytes)
	revBytes := make([]byte, size)
	for i := 0; i < size; i++ {
		revBytes[i] = bytes[size-1-i]
	}
	return string(revBytes)
}
