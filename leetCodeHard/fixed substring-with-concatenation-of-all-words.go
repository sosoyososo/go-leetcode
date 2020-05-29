package leetCodeHard

import (
	"fmt"
	"sort"
)

func findSubstring(s string, words []string) sort.IntSlice {
	if len(s) == 0 || len(words) == 0 {
		return sort.IntSlice{}
	}

	doesMatchFromIndex := func(index int, key string) bool {
		for i, v := range key {
			if index+i < len(s) && v == rune(s[index+i]) {
				continue
			}
			return false
		}
		return true
	}

	indexesMap := map[string]sort.IntSlice{}
	for _, word := range words {
		//避免匹配词语列表中重复词语出现
		if len(indexesMap[word]) > 0 {
			break
		}
		for j := 0; j < len(s); j++ {
			if doesMatchFromIndex(j, word) {
				list := indexesMap[word]
				list = append(list, j)
				indexesMap[word] = list
			}
		}
	}

	numsOfWord := len(words)
	if numsOfWord == 1 {
		return indexesMap[words[0]]
	}

	totalLen := 1
	indexesListOrig := []sort.IntSlice{}
	for _, word := range words {
		indexes := indexesMap[word]
		totalLen *= len(indexes)
		indexesListOrig = append(indexesListOrig, indexes)
	}
	if totalLen <= 0 {
		return []int{}
	}

	indexesList := []sort.IntSlice{}
	for i := 0; i < totalLen; i++ {
		indexesList = append(indexesList, make(sort.IntSlice, numsOfWord))
	}

	step := totalLen
	for i, indexes := range indexesListOrig {
		step /= len(indexes)
		for j := 0; j < totalLen; j++ {
			jIndex := j / step % len(indexes)
			indexesList[j][i] = indexes[jIndex]
		}
	}

	indexes := []int{}

	insertIndexIfCan := func(list sort.IntSlice) {
		if len(list) == 0 {
			return
		}
		for i := 0; i < len(list)-1; i++ {
			if list[i+1]-list[i] != len(words[0]) {
				return
			}
		}
		for _, v := range indexes {
			if v == list[0] {
				return
			}
		}
		indexes = append(indexes, list[0])
	}
	for _, list := range indexesList {
		sort.Sort(list)
		insertIndexIfCan(list)
	}

	fmt.Println(indexes)
	return indexes
}
