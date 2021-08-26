package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/validate-binary-search-tree/
 */

// IsValidBSTV1 递归方式
func IsValidBSTV1(root *TreeNode) bool {
	valid, _, _ := checkBST(root)
	return valid
}

func checkBST(root *TreeNode) (valid bool, min int, max int) {
	if root == nil {
		return
	}
	valid, min, max = true, root.Val, root.Val
	if root.Left != nil {
		tmpValid, tmpMin, tmpMax := checkBST(root.Left)
		if !tmpValid || tmpMax >= root.Val {
			return false, 0, 0
		}
		if tmpMin < min {
			min = tmpMin
		}
	}
	if root.Right != nil {
		tmpValid, tmpMin, tmpMax := checkBST(root.Right)
		if !tmpValid || tmpMin <= root.Val {
			return false, 0, 0
		}
		if tmpMax > max {
			max = tmpMax
		}
	}
	return
}

// IsValidBSTV2 性质：二叉搜索树的中序遍历结果为递增序列
func IsValidBSTV2(root *TreeNode) bool {
	list := make([]int, 0)
	var inOrderTraversal func(*TreeNode)
	inOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrderTraversal(node.Left)
		list = append(list, node.Val)
		inOrderTraversal(node.Right)
	}
	inOrderTraversal(root)
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}
	return true
}
