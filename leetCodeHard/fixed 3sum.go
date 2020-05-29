package leetCodeHard

import "sort"

func threeSum(nums []int) [][]int {
	var sliceNums sort.IntSlice
	sliceNums = nums
	sort.Sort(sliceNums)

	ret := [][]int{}
	match := func(list []int) bool {
		for _, v := range ret {
			if list[0] == v[0] && list[1] == v[1] {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(sliceNums); i++ {
		for j := i + 1; j < len(sliceNums); j++ {
			if i == j {
				continue
			}
			s := sliceNums[i] + sliceNums[j]
			if findNum(sliceNums[j+1:], 0-s) {
				tmpList := sort.IntSlice{sliceNums[i], sliceNums[j], 0 - s}
				sort.Sort(tmpList)
				if !match(tmpList) {
					ret = append(ret, tmpList)
				}
			}
		}
	}

	return ret
}

func findNum(list []int, num int) bool {
	l := len(list)
	i := l / 2
	if l == 0 {
		return false
	}
	if list[i] == num {
		return true
	} else if list[i] < num {
		return findNum(list[i+1:], num)
	} else {
		return findNum(list[:i], num)
	}
}
