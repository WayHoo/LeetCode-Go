package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	curLevelCnt, nextLevelCnt := 1, 0
	nums := make([]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextLevelCnt++
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextLevelCnt++
		}
		nums = append(nums, cur.Val)
		if curLevelCnt--; curLevelCnt == 0 {
			ans = append(ans, append([]int{}, nums...))
			nums = nil
			curLevelCnt = nextLevelCnt
			nextLevelCnt = 0
		}
		queue = queue[1:]
	}
	return ans
}
