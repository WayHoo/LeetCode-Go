package problems

/*
 LeetCode: https://leetcode-cn.com/problems/permutations-ii/
*/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tailA, tailB := headA, headB
	lenA, lenB := 1, 1
	for ; tailA.Next != nil; tailA, lenA = tailA.Next, lenA+1 {}
	for ; tailB.Next != nil; tailB, lenB = tailB.Next, lenB+1 {}
	if tailA != tailB {
		return nil
	}
	tmpHeadA, tmpHeadB := headA, headB
	for i := 0; i < lenA-lenB; tmpHeadA, i = tmpHeadA.Next, i+1 {}
	for i := 0; i < lenB-lenA; tmpHeadB, i = tmpHeadB.Next, i+1 {}
	for tmpHeadA != nil && tmpHeadB != nil {
		if tmpHeadA == tmpHeadB {
			return tmpHeadA
		}
		tmpHeadA, tmpHeadB = tmpHeadA.Next, tmpHeadB.Next
	}
	return nil
}

// GetIntersectionNodeV2 骚操作方式
func GetIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}
