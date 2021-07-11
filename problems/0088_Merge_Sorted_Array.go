package problems

/*
 LeetCode: https://leetcode-cn.com/problems/merge-sorted-array/
*/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, j, k := m-1, n-1, m+n-1
	for ; j >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
