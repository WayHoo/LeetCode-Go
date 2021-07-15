package main

import (
	"LeetCode-Go/algorithm"
	"LeetCode-Go/problems"
	"fmt"
)

func main() {
	res := problems.IntegerBreak(58)
	fmt.Println(res)

	//testSort()
	//testLRUCache()
	//testUnionFindSet()
}

func testSort() {
	T := []int{5, 1, 2, 9, 7, 4, 8, 3, 10, 6}
	var T1, T2, T3, T4, T5 []int
	T1 = append(T1, T...)
	algorithm.Bubble(T1)
	fmt.Println("Bubble sort:", T1)
	T2 = append(T2, T...)
	algorithm.Insertion(T2)
	fmt.Println("Insertion sort:", T2)
	T3 = append(T3, T...)
	algorithm.Shell(T3)
	fmt.Println("Shell sort:", T3)
	T4 = append(T4, T...)
	algorithm.Up2DownMergeSort(T4)
	fmt.Println("Up2DownMergeSort sort:", T4)
	T5 = append(T5, T...)
	algorithm.Down2UpMergeSort(T5)
	fmt.Println("Down2UpMergeSort sort:", T5)
}

func testLRUCache() {
	lru := problems.Constructor(2)
	lru.Watch()
	lru.Put(1, 1)
	lru.Watch()
	lru.Put(2, 2)
	lru.Watch()
	fmt.Println(lru.Get(1))
	lru.Watch()
	lru.Put(3, 3)
	lru.Watch()
	fmt.Println(lru.Get(2))
	lru.Watch()
	lru.Put(4, 4)
	lru.Watch()
	fmt.Println(lru.Get(1))
	lru.Watch()
	fmt.Println(lru.Get(3))
	lru.Watch()
	fmt.Println(lru.Get(4))
	lru.Watch()
}

func testUnionFindSet() {
	uf := algorithm.UFSetConstructor(8)
	uf.Union(3, 2)
	uf.Union(2, 4)
	uf.Union(2, 5)
	uf.Union(1, 6)
	uf.Union(6, 7)
	uf.Union(4, 6)
	_ = uf.Find(1)
	_ = uf.Find(7)
}
