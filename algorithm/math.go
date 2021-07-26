package algorithm

// GCD 最大公约数（Greatest Common Divisor），辗转相除法
func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    if a < b {
		a, b = b, a
	}
	return GCD(b, a%b)
}

// LCM 最小公倍数为两数的乘积除以最大公约数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
