package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/swap-nodes-in-pairs/
 */

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, nnext := head.Next, head.Next.Next
	head.Next.Next = head
	head.Next = SwapPairs(nnext)
	return newHead
}
