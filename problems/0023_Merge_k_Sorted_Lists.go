package problems

import "sort"

/**
 * LeetCode: https://leetcode-cn.com/problems/merge-k-sorted-lists/
 */

func MergeKLists(lists []*ListNode) *ListNode {
	total := 0
	for _, node := range lists {
		for ; node != nil; node = node.Next {
			total++
		}
	}
	if total == 0 {
		return nil
	}
	nodeList, idx := make([]*ListNode, total), 0
	for _, node := range lists {
		for ; node != nil; node = node.Next {
			nodeList[idx] = node
			idx++
		}
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].Val <= nodeList[j].Val
	})
	for i := 0; i < total-1; i++ {
		nodeList[i].Next = nodeList[i+1]
	}
	nodeList[total-1].Next = nil
	return nodeList[0]
}
