package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 */

// MaxDepth 递归方法
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// MaxDepthV2 层次遍历方法
func MaxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level, curLevelNum, nextLevelNum := 0, 1, 0
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextLevelNum++
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextLevelNum++
		}
		if curLevelNum--; curLevelNum == 0 {
			level++
			curLevelNum, nextLevelNum = nextLevelNum, 0
		}
		queue = queue[1:]
	}
	return level
}
