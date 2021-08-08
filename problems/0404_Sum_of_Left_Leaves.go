package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/sum-of-left-leaves/
 */

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root.Left) {
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}
	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}

func isLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func SumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		cur := queue[0]
		if left := cur.Left; left != nil {
			if left.Left == nil && left.Right == nil {
				sum += left.Val
			} else {
				queue = append(queue, left)
			}
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		queue = queue[1:]
	}
	return sum
}
