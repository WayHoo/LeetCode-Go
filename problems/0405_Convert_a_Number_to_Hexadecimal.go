package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/
 */

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}
	char := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var bytes []byte
	for num != 0 {
		n := num % 16
		bytes = append(bytes, char[n])
		num /= 16
	}
	size := len(bytes)
	revBytes := make([]byte, size)
	for i := 0; i < size; i++ {
		revBytes[i] = bytes[size-1-i]
	}
	return string(revBytes)
}
