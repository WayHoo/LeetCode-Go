package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 */

func FindKthLargest(nums []int, k int) int {
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    medium := func(l, r int) {
        if l >= r {
            return
        }
        m := l + (r - l)/2
        if nums[l] > nums[m] {
            swap(l, m)
        }
        if nums[l] > nums[r] {
            swap(l, r)
        }
        if nums[m] > nums[r] {
            swap(m, r)
        }
        swap(m, r-1)
    }
    var quickSort func(l, r int) int
    quickSort = func(l, r int) int {
        if l >= r {
            return nums[l]
        }
        medium(l, r)
        i, j, pivot := l, r-1, nums[r-1]
        for i < j {
            for i++; nums[i] < pivot; i++ {}
            for j--; nums[j] > pivot; j-- {}
            if i < j {
                swap(i, j)
            }
        }
        swap(i, r-1)
        if i == len(nums)-k {
            return nums[i]
        } else if i > len(nums) - k {
            return quickSort(l, i-1)
        } else {
            return quickSort(i+1, r)
        }
    }
    return quickSort(0, len(nums)-1)
}
