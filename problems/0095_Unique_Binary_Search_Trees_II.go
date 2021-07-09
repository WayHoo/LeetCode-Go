package problems

/*
 LeetCode: https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
*/

func GenerateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	return formTrees(1, n)
}

func formTrees(m, n int) []*TreeNode {
	if m > n {
		return []*TreeNode{nil}
	}
	trees := make([]*TreeNode, 0)
	for i := m; i <= n; i++ {
		leftTrees := formTrees(m, i-1)
		rightTrees := formTrees(i+1, n)
		for _, l := range leftTrees {
			for _, r := range rightTrees {
				tree := &TreeNode{i, l, r}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}
