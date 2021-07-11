package problems

/*
 LeetCode: https://leetcode-cn.com/problems/word-ladder/
*/

func LadderLength(beginWord string, endWord string, wordList []string) int {
	queue := []int{-1}
	visited := make([]bool, len(wordList))
	level, curLevelNum, nextLevelNum := 1, 1, 0
	for len(queue) > 0 {
		idx := queue[0]
		str := beginWord
		if idx >= 0 {
			str = wordList[idx]
		}
		for i, word := range wordList {
			if !visited[i] && canConvert(str, word) {
				if word == endWord {
					return level + 1
				}
				visited[i] = true
				queue = append(queue, i)
				nextLevelNum++
			}
		}
		if curLevelNum--; curLevelNum == 0 {
			level++
			curLevelNum, nextLevelNum = nextLevelNum, 0
		}
		queue = queue[1:]
	}
	return 0
}

func canConvert(source, target string) bool {
	diff := 0
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			if diff++; diff > 1 {
				return false
			}
		}
	}
	return true
}
