package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/base-7/
 */

func ConvertToBase7(num int) string {
    if num == 0 {
        return "0"
    }
    posiNum := num
    if num < 0 {
        posiNum = -num
    }
    var bytes []byte
    for posiNum > 0 {
        n := posiNum % 7
        bytes = append(bytes, byte('0' + n))
        posiNum /= 7
    }
    size := len(bytes)
    revBytes := make([]byte, size)
    for i := 0; i < size; i++ {
        revBytes[i] = bytes[size-1-i]
    }
    str := string(revBytes)
    if num < 0 {
        str = "-" + str
    }
    return str
}
