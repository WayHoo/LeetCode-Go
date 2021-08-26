package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 */

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return node
	}
	idx := findIdx(preorder[0], inorder)
	leftInorder := inorder[:idx]
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightInorder := inorder[idx+1:]
	rightPreorder := preorder[1+len(leftInorder):]
	node.Left = BuildTree(leftPreorder, leftInorder)
	node.Right = BuildTree(rightPreorder, rightInorder)
	return node
}

func findIdx(target int, inorder []int) int {
	for i, num := range inorder {
		if num == target {
			return i
		}
	}
	return -1
}
