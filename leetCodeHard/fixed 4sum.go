package leetCodeHard

import "sort"

func fourSum(nums []int, target int) [][]int {
	var sliceNums sort.IntSlice
	sliceNums = nums
	sort.Sort(sliceNums)

	ret := [][]int{}
	match := func(list []int) bool {
		for _, v := range ret {
			if list[0] == v[0] && list[1] == v[1] && list[2] == v[2] {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(sliceNums); i++ {
		for j := i + 1; j < len(sliceNums); j++ {
			for k := j + 1; k < len(sliceNums); k++ {
				s := sliceNums[i] + sliceNums[j] + sliceNums[k]
				if findNum(sliceNums[k+1:], target-s) {
					tmpList := sort.IntSlice{sliceNums[i], sliceNums[j], sliceNums[k], target - s}
					sort.Sort(tmpList)
					if !match(tmpList) {
						ret = append(ret, tmpList)
					}
				}
			}

		}
	}

	return ret
}
