package problems

/*
 LeetCode: https://leetcode-cn.com/problems/reverse-linked-list/
*/

func ReverseList(head *ListNode) *ListNode {
	var node *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		node = &ListNode{cur.Val, node}
	}
	return node
}
