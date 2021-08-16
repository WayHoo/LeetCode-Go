package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
 */

func InorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var ret []int
    stk := []*TreeNode{root}
    vis := make(map[*TreeNode]int)
    for len(stk) > 0 {
        top := stk[len(stk)-1]
        vis[top]++
        if vis[top] == 1 && top.Left != nil {
            stk = append(stk, top.Left)
        } else if vis[top] == 2 {
            ret = append(ret, top.Val)
            if top.Right != nil {
                stk = append(stk, top.Right)
            }
        } else if vis[top] == 3 {
            stk = stk[:len(stk)-1]
        }
    }
    return ret
}
