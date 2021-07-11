package problems

import "strconv"

/**
 * LeetCode: https://leetcode-cn.com/problems/restore-ip-addresses/
 */

func RestoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	res := backTrack(s, 4)
	return res
}

func backTrack(s string, cnt int) (res []string) {
	if s != "" && cnt <= 0 {
		return
	}
	if s == "" {
		if cnt == 0 {
			res = append(res, "")
		}
		return
	}
	for i := 0; i < 3 && i < len(s); i++ {
		num, _ := strconv.ParseInt(s[:i+1], 10, 64)
		if num <= 255 {
			right := backTrack(s[i+1:], cnt-1)
			for _, str := range right {
				tmp := s[:i+1]
				if str != "" {
					tmp = tmp + "." + str
				}
				res = append(res, tmp)
			}
		} else {
			break
		}
		if i == 0 && s[i] == '0' {
			break
		}
	}
	return
}
