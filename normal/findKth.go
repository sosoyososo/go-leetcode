package normal

import (
	"errors"
)

type item struct {
	value int
	index int
}

func FindKth(arr []int, k int, canRepeat bool) (int, int) {
	if len(arr) < k {
		panic(errors.New("数组长度不够"))
	}
	if k <= 0 {
		panic(errors.New("参数错误"))
	}

	// findMinIndex := func(arr []item) int {
	// 	minIndex := 0
	// 	for j := minIndex; j < len(arr); j++ {
	// 		if arr[j].value < arr[minIndex].value {
	// 			minIndex = j
	// 		}
	// 	}
	// 	return minIndex
	// }

	findRepeat := func(arr []item, ele int) bool {
		for i := 0; i < len(arr); i++ {
			if arr[i].value == ele {
				return true
			}
		}
		return false
	}

	repeatMinElem := func(arr []item, ele, index int) []item {
		ret := []item{}
		hasInsert := false
		for i := 1; i < len(arr); i++ {
			if arr[i].value < ele {
				ret = append(ret, arr[i])
			} else {
				hasInsert = true
				ret = append(ret, item{value: ele, index: index})
				ret = append(ret, arr[i])
			}
		}
		if !hasInsert {
			ret = append(ret, item{value: ele, index: index})
		}

		return ret
	}

	ret := []item{}
	for i := 0; i < len(arr); i++ {
		if !canRepeat {
			repeat := findRepeat(ret, arr[i])
			if repeat {
				continue
			}
		}
		if len(ret) < k {
			ret = append(ret, item{value: arr[i], index: i})
			continue
		}

		if ret[0].value < arr[i] {
			ret = repeatMinElem(ret, arr[i], i)
		}
	}

	ele := ret[0]
	return ele.value, ele.index
}

// func TestKth() {
// 	for i := 0; i < 100; i++ {
// 		k := 5
// 		arr := sort.GenerateRanomArr(100)
// 		ele := FindKth(arr, k)
// 		sort.BubbleSort(&arr)
// 		if ele != arr[len(arr)-5] {
// 			fmt.Println(arr)
// 			fmt.Println(ele)
// 			panic(errors.New("Kth 查找错误"))
// 		}
// 	}
// }
