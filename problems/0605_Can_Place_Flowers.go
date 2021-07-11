package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
 */

func CanPlaceFlowers(flowerbed []int, n int) bool {
	N := len(flowerbed)
	if N == 0 {
		return false
	}
	place := 0
	for i := 0; i < N; i++ {
		if flowerbed[i] == 0 && (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 >= N || flowerbed[i+1] == 0) {
			place++
			i++
		}
	}
	return place >= n
}
