package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/add-two-numbers-ii/
 */

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := make([]int, 0), make([]int, 0)
	for node := l1; node != nil; node = node.Next {
		num1 = append(num1, node.Val)
	}
	for node := l2; node != nil; node = node.Next {
		num2 = append(num2, node.Val)
	}
	sum, c := make([]int, 0), 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if i >= 0 {
			a = num1[i]
		}
		if j >= 0 {
			b = num2[j]
		}
		if t := a + b + c; t >= 10 {
			c = 1
			sum = append(sum, t%10)
		} else {
			c = 0
			sum = append(sum, t)
		}
	}
	if c > 0 {
		sum = append(sum, c)
	}
	head := &ListNode{}
	pre := head
	for i := len(sum) - 1; i >= 0; i-- {
		node := &ListNode{Val: sum[i]}
		pre.Next = node
		pre = node
	}
	return head.Next
}
