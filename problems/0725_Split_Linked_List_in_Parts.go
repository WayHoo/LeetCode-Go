package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/split-linked-list-in-parts/
 */

func SplitListToParts(head *ListNode, k int) []*ListNode {
	ans, idx := make([]*ListNode, k), 0
	cnt := 0
	for node := head; node != nil; node = node.Next {
		cnt++
	}
	cur := head
	size, mod := cnt/k, cnt%k
	for i := 0; cur != nil && i < k; i++ {
		curSize := size
		if mod > 0 {
			curSize++
			mod--
		}
		ans[idx] = cur
		idx++
		for j := 0; j < curSize-1; j++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	return ans
}
