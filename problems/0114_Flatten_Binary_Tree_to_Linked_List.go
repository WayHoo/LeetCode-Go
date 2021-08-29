package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
 */

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var subFlatten func(*TreeNode) *TreeNode
	subFlatten = func(tree *TreeNode) *TreeNode {
		node := tree
		right := node.Right
		if node.Left != nil {
			node.Right = subFlatten(node.Left)
			node.Left = nil
			for node.Right != nil {
				node = node.Right
			}
		}
		if right != nil {
			node.Right = subFlatten(right)
		}
		return tree
	}
	subFlatten(root)
}
