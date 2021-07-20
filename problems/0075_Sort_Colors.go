package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/sort-colors/
 */

func SortColors(nums []int)  {
    lt, gt, i, pivot := 0, len(nums)-1, 0, 1
    for i <= gt {
        if nums[i] < pivot {
            nums[i], nums[lt] = nums[lt], nums[i]
            i++
            lt++
        } else if nums[i] > pivot {
            nums[i], nums[gt] = nums[gt], nums[i]
            gt--
        } else {
            i++
        }
    }
    return
}
