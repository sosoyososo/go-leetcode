package leetCodeHard

import (
	"fmt"
	"sort"
)

func Test() {
	list := []int{2, 3, 1}
	nextPermutation(list)
	fmt.Println(list)
	// findSubstring("aaa", []string{"a", "b"})
	// test()
	// testSlice()
}

func testSlice() {
	list := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 3, 4}
	slice := list[:]
	slice = append(slice, 5)
	slice2 := slice
	slice = append(slice, 5)
	slice[0] = 0
	fmt.Println(slice)
	fmt.Println(slice2)
}

func test() {
	wordsIndexesMap := map[string]sort.IntSlice{"ab": sort.IntSlice{0, 3, 5, 8}, "ba": sort.IntSlice{1, 4, 7, 9}}
	words := []string{"ab", "ba", "ab", "ba"}

	indexesRet := []sort.IntSlice{}
	for _, word := range words {
		indexList := wordsIndexesMap[word]

		tmpRet := []sort.IntSlice{}
		for _, index := range indexList {
			if len(indexesRet) > 0 {
				for _, indexes := range indexesRet {
					indexes = append(indexes, index)
					tmpRet = append(tmpRet, indexes)
				}
			} else {
				tmpRet = append(tmpRet, sort.IntSlice{index})
			}
		}
		indexesRet = tmpRet
	}

	for _, ll := range indexesRet {
		sort.Sort(ll)
		fmt.Println(ll)
	}
}
