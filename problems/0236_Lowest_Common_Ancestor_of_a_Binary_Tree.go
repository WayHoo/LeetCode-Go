package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 */

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var commonAncestor func(*TreeNode) (int, *TreeNode)
	commonAncestor = func(node *TreeNode) (int, *TreeNode) {
		count := 0
		if node == p || node == q {
			count++
		}
		if node.Left != nil {
			cnt, ancestor := commonAncestor(node.Left)
			if cnt == 2 {
				return cnt, ancestor
			}
			count += cnt
		}
		if node.Right != nil {
			cnt, ancestor := commonAncestor(node.Right)
			if cnt == 2 {
				return cnt, ancestor
			}
			count += cnt
		}
		return count, node
	}
	_, ancestor := commonAncestor(root)
	return ancestor
}
