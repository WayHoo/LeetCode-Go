package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/balanced-binary-tree/
 */

func IsBalanced(root *TreeNode) bool {
	var checkBalance func(root *TreeNode) (bool, int)
	checkBalance = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}
		lc, lh := checkBalance(root.Left)
		rc, rh := checkBalance(root.Right)
		if lc && rc && lh-rh >= -1 && lh-rh <= 1 {
			h := lh
			if rh > lh {
				h = rh
			}
			return true, h+1
		} else {
			return false, 0
		}
	}
	check, _ := checkBalance(root)
	return check
}
