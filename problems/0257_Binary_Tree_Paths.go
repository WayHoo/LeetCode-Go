package problems

import "fmt"

/**
 * LeetCode: https://leetcode-cn.com/problems/binary-tree-paths/
 */

func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	numStr := fmt.Sprintf("%d", root.Val)
	res := make([]string, 0)
	list := append(BinaryTreePaths(root.Left), BinaryTreePaths(root.Right)...)
	if len(list) == 0 {
		res = append(res, numStr)
	}
	for _, val := range list {
		res = append(res, numStr + "->" + val)
	}
	return res
}
