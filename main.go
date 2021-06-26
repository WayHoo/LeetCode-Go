package main

import (
	"fmt"

	"LeetCode-Go/problems"
)

func main() {
	res := problems.FindLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"})
	fmt.Println(res)
}
