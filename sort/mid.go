package sort

import (
	"fmt"
	"strings"
)

/**
 * 中值排序先找到中值元素，与中位元素交换位置，然后大的放右边，小的放左边；
 * 再分别对两边执行中值排序
 * 边界处有可能过界，结束后需要修正
 * 按照Kth方式，可能出现无效分割，4 0 4 4 的情况下，
 * 按照4进行分割，就会都出现在左，无效分割可能是因为等值过多，也可能全都是等值(⊙o⊙)…
 * 眼高手低的结果，一个边界问题处理了三天
 */
func MidSort(arr *[]int) {
	if !needSort(arr) {
		return
	}
	quickSort(arr, 0, len(*arr)-1)
}

func quickSort(arr *[]int, left, right int) {
	if right <= left {
		return
	}
	index := quickSplit2(arr, left, right)
	quickSort(arr, left, index-1)
	quickSort(arr, index+1, right)
}

/**
 * left of store is less,
 * right of store and left of i is bigger,
 * and right of i is not tested
 */
func quickSplit(arr *[]int, left, right int) int {
	swap(arr, right, left)
	store := left
	for i := left; i < right; i++ {
		if (*arr)[i] <= (*arr)[right] {
			swap(arr, i, store)
			store++
		}
	}
	swap(arr, store, right)
	return store
}

func quickSplit2(arr *[]int, leftIndex, rightIndex int) int {
	left := leftIndex
	right := rightIndex
	v := (*arr)[leftIndex]
	for {
		if (*arr)[left] > v && (*arr)[right] <= v {
			swap(arr, left, right)
			continue
		}
		if left == right {
			break
		}
		if (*arr)[left] <= v {
			left++
			continue
		}
		if (*arr)[right] > v {
			right--
			continue
		}
	}
	if (*arr)[left] > v && left > leftIndex {
		left--
	}
	if (*arr)[right] <= v && right < rightIndex {
		right++
	}
	swap(arr, leftIndex, left)
	return left
}

func arrayDesc(arr []int) string {
	ret := []string{}
	for _, v := range arr {
		ret = append(ret, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[]int{%v}", strings.Join(ret, ","))
}
