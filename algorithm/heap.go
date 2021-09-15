package algorithm

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// 这里决定是大顶堆还是小顶堆，现在是小顶堆
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	// 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func TestHeap() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)
	heap.Push(h, 6)
	for len(*h) > 0 {
		fmt.Println(heap.Pop(h))
	}
}
