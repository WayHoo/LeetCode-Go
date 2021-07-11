package problems

import (
	"bytes"
)

/*
 LeetCode: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
*/

func LetterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	digit := digits[0]
	chs := digit2chs(digit)
	for i := range chs {
		result := LetterCombinations(digits[1:])
		if len(result) == 0 {
			res = append(res, chs[i:i+1])
		}
		for _, str := range result {
			var buffer bytes.Buffer
			buffer.WriteString(chs[i:i+1])
			buffer.WriteString(str)
			res = append(res, buffer.String())
		}
	}
	return res
}

func digit2chs(digit uint8) string {
	switch digit {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}
