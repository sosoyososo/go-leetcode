package sort

import (
	"errors"
	"fmt"
)

func QuickSort(arr *[]int) {
	if !needSort(arr) {
		return
	}
	relQickSort(arr, 0, len(*arr)-1)
}

func relQickSort(arr *[]int, left, right int) {
	if right-left < 1 {
		return
	}
	mid := pickMid(arr, left, right) //选择mid不当会导致死循环
	index := split(arr, mid, left, right)
	relQickSort(arr, left, index) //等于mid的放在小头
	relQickSort(arr, index+1, right)
}

func pickMid(arr *[]int, left, right int) int {
	return ((*arr)[left] + (*arr)[right]) / 2
}

func split(arr *[]int, mid, leftIndx, rightIndx int) int { //不同位置处理//左侧//右侧//
	left := leftIndx
	right := rightIndx
	for {
		if right-left <= 1 {
			break
		}
		if (*arr)[left] <= mid {
			left++
			continue
		}
		if (*arr)[right] > mid {
			right--
			continue
		}
		swap(arr, left, right)
	}

	if right == rightIndx {
		fmt.Println(*arr)
		fmt.Printf("left %v right %v\n", leftIndx, rightIndx)
		panic(errors.New("位置错误1"))
	} else if leftIndx == left {
		fmt.Println(arr)
		fmt.Printf("left %v right %v", leftIndx, rightIndx)
		panic(errors.New("位置错误2"))
	} else {
		return right
	}
}
