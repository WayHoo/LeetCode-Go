package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/top-k-frequent-elements/
 */

func TopKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    buckets := make([][]int, len(nums)+1)
    for num, freq := range freqMap {
        buckets[freq] = append(buckets[freq], num)
    }
    ret := make([]int, 0, k)
    for i, cnt := len(nums), 0; i > 0 && cnt <= k; i-- {
        N := len(buckets[i])
        if N == 0 {
            continue
        }
        cnt += N
        ret = append(ret, buckets[i]...)
    }
    return ret[:k]
}
