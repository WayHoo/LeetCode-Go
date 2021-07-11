package problems

/*
 LeetCode: https://leetcode-cn.com/problems/diameter-of-binary-tree/
*/

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := InvertTree(root.Left)
	r := InvertTree(root.Right)
	root.Left = r
	root.Right = l
	return root
}
