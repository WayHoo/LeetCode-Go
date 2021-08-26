package main

import (
    "bufio"
    "fmt"
    "os"

    "LeetCode-Go/algorithm"
    "LeetCode-Go/concurrency"
    "LeetCode-Go/problems"
)

func main() {
    res := problems.SubsetsBits([]int{1, 2, 3})
    fmt.Println(res)

    //testSort()
    //testLRUCache()
    //testUnionFindSet()
    //subMain()
    //testConcurrency()
}

func testSort() {
	T := []int{5, 1, 2, 9, 7, 4, 8, 3, 10, 6, 6, 2, 4, 1, 5, 3, 9, 8, 7, 10}
	var T1, T2, T3, T4, T5, T6, T7, T8 []int
	T1 = append(T1, T...)
	algorithm.BubbleSort(T1)
	fmt.Println("BubbleSort:", T1)
	T2 = append(T2, T...)
	algorithm.InsertionSort(T2)
	fmt.Println("InsertionSort:", T2)
	T3 = append(T3, T...)
	algorithm.ShellSort(T3)
	fmt.Println("ShellSort:", T3)
	T4 = append(T4, T...)
	algorithm.Up2DownMergeSort(T4)
	fmt.Println("Up2DownMergeSort:", T4)
	T5 = append(T5, T...)
	algorithm.Down2UpMergeSort(T5)
	fmt.Println("Down2UpMergeSort:", T5)
	T6 = append(T6, T...)
	algorithm.QuickSort(T6)
	fmt.Println("QuickSort:", T6)
	T7 = append(T7, T...)
	algorithm.ThreeWayQuickSort(T7)
	fmt.Println("ThreeWayQuickSort:", T7)
	T8 = append(T8, T...)
	algorithm.HeapSort(T8)
	fmt.Println("HeapSort:", T8)
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

func subMain() {
	name, age, salary := "", 0, 0.0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%s%d%f", &name, &age, &salary)
		if err != nil {
			break
		}
		fmt.Println(name, age, salary)
	}
}

func testConcurrency() {
	concurrency.Quiz02()
}
