package problems

import "sort"

/*
 LeetCode: https://leetcode-cn.com/problems/queue-reconstruction-by-height/
*/

func ReconstructQueue(people [][]int) [][]int {
	N := len(people)
	if N == 0 {
		return nil
	}
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	queue := make([][]int, 0, N)
	for _, person := range people {
		idx := person[1]
		queue = append(queue[:idx], append([][]int{person}, queue[idx:]...)...)
	}
	return queue
}
