package problems

/*
 LeetCode: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 解题思路：实际上就是获取链表倒数第 n+1 个结点，需注意如果删除的是头结点的边界情况。
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	l, r := head, head
	i := 0
	for ; i <= n && r != nil; i++ {
		r = r.Next
	}
	if i <= n {
		return head.Next
	}
	for ; r != nil; l, r = l.Next, r.Next {}
	l.Next = l.Next.Next
	return head
}
