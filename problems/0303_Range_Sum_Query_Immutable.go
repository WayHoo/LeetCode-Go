package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/range-sum-query-immutable/
 */

type NumArray struct {
    SubSums []int
}

func NumArrayConstructor(nums []int) NumArray {
    arr := NumArray{make([]int, len(nums)+1)}
    for i := 1; i <= len(nums); i++ {
        arr.SubSums[i] = arr.SubSums[i-1] + nums[i-1]
    }
    return arr
}

func (this *NumArray) SumRange(left int, right int) int {
    return this.SubSums[right+1] - this.SubSums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
