package problems

/*
 LeetCode: https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
*/

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	h1, h2, h := l1, l2, head
	for ; h1 != nil && h2 != nil; h = h.Next {
		if h1.Val <= h2.Val {
			h.Next = h1
			h1 = h1.Next
		} else {
			h.Next = h2
			h2 = h2.Next
		}
	}
	for ;h1 != nil; h = h.Next {
		h.Next = h1
		h1 = h1.Next
	}
	for ;h2 != nil; h = h.Next {
		h.Next = h2
		h2 = h2.Next
	}
	return head.Next
}
