package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/linked-list-cycle/
 */

// 时间复杂度和空间复杂度均为 O(n)
func HasCycle1(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if nodeMap[node] {
			return true
		} else {
			nodeMap[node] = true
			node = node.Next
		}
	}
	return false
}

// 判断是否存在环的问题，可设置移动速度不同的两个指针，转换成追击相遇问题。
// 时间复杂度 O(n)，空间复杂度 O(1)
func HasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	n1, n2 := head, head.Next
	for n2 != nil && n2.Next != nil {
		if n1 == n2 {
			return true
		} else {
			n1 = n1.Next
			n2 = n2.Next.Next
		}
	}
	return false
}
