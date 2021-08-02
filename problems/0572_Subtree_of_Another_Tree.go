package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/subtree-of-another-tree/
 */

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return matchTree(root, subRoot) || IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func matchTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}
	if root.Val != subRoot.Val {
		return false
	}
	return matchTree(root.Left, subRoot.Left) && matchTree(root.Right, subRoot.Right)
}
