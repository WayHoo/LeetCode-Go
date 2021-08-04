package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/palindrome-linked-list/
 */

func IsPalindromeLinkedList(head *ListNode) bool {
    if head == nil {
        return false
    }
    nums := make([]int, 0)
    for node := head; node != nil; node = node.Next {
        nums = append(nums, node.Val)
    }
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        if nums[i] != nums[j] {
            return false
        }
    }
    return true
}
