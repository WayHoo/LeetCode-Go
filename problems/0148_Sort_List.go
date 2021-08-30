package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/sort-list/
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	leftPart, rightPart := cutList(head)
	leftPart = sortList(leftPart)
	rightPart = sortList(rightPart)
	return mergeList(leftPart, rightPart)
}

func cutList(head *ListNode) (leftPart, rightPart *ListNode) {
	if head == nil {
		return
	}
	cnt := 0
	for node := head; node != nil; node = node.Next {
		cnt++
	}
	node := head
	for i := 0; i < cnt/2-1; node, i = node.Next, i+1 {
	}
	rightPart = node.Next
	node.Next = nil
	leftPart = head
	return
}

func mergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head := &ListNode{}
	node1, node2, node := head1, head2, head
	for ; node1 != nil && node2 != nil; node = node.Next {
		if node1.Val <= node2.Val {
			node.Next = node1
			node1 = node1.Next
		} else {
			node.Next = node2
			node2 = node2.Next
		}
	}
	for ; node1 != nil; node, node1 = node.Next, node1.Next {
		node.Next = node1
	}
	for ; node2 != nil; node, node2 = node.Next, node2.Next {
		node.Next = node2
	}
	return head.Next
}
