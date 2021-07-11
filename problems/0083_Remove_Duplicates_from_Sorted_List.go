package problems

/*
 LeetCode: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
*/

func DeleteDuplicates(head *ListNode) *ListNode {
	for h := head; h != nil; {
		if h.Next == nil {
			break
		}
		if h.Val == h.Next.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return head
}
