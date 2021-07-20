package problems

import "sort"

/**
 * LeetCode: https://leetcode-cn.com/problems/sort-characters-by-frequency/
 */

func FrequencySort(s string) string {
    type CharFreq struct {
        Ch   byte
        Freq int
    }
    freqMap := make(map[byte]int)
    for _, ch := range s {
        freqMap[byte(ch)]++
    }
    freqList := make([]CharFreq, 0, len(freqMap))
    for ch, freq := range freqMap {
        freqList = append(freqList, CharFreq{ch, freq})
    }
    sort.Slice(freqList, func(i, j int) bool {
        if freqList[i].Freq >= freqList[j].Freq {
            return true
        }
        return false
    })
    list := make([]byte, len(s))
    idx := 0
    for _, val := range freqList {
        for i := 0; i < val.Freq; i++ {
            list[idx] = val.Ch
            idx++
        }
    }
    return string(list)
}
