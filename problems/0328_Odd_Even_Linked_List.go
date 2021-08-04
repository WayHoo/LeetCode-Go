package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/odd-even-linked-list/
 */

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
