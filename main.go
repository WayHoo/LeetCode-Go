package main

import (
	"fmt"

	"LeetCode-Go/problems"
	mySort "LeetCode-Go/sort"
)

func main() {
	res := problems.Permute([]int{1, 2, 3})
	fmt.Println(res)
	//testSort()
}

func testSort() {
	T := []int{5, 1, 2, 9, 7, 4, 8, 3, 10, 6}
	var T1, T2, T3, T4, T5 []int
	T1 = append(T1, T...)
	mySort.Bubble(T1)
	fmt.Println("Bubble sort:", T1)
	T2 = append(T2, T...)
	mySort.Insertion(T2)
	fmt.Println("Insertion sort:", T2)
	T3 = append(T3, T...)
	mySort.Shell(T3)
	fmt.Println("Shell sort:", T3)
	T4 = append(T4, T...)
	mySort.Up2DownMergeSort(T4)
	fmt.Println("Up2DownMergeSort sort:", T4)
	T5 = append(T5, T...)
	mySort.Down2UpMergeSort(T5)
	fmt.Println("Down2UpMergeSort sort:", T5)
}
