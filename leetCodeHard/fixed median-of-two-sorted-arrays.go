package leetCodeHard

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func midOfTwoSortedArr(arr1 []int, arr2 []int) float32 {
	l1 := len(arr1)
	l2 := len(arr2)

	midFunc := func(arr []int) float32 {
		l := len(arr)
		if l%2 == 0 {
			return (float32(arr[l/2]-1) + float32(arr[l/2])) / 2.0
		} else {
			return float32(arr[l/2])
		}
	}

	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		return midFunc(arr2)
	} else if l2 == 0 {
		return midFunc(arr1)
	}

	ret := []int{}
	index1 := 0
	index2 := 0
	stopLen := (l1+l2)/2 + 1
	for {
		if len(ret) == stopLen {
			break
		}
		if index1 < l1 && (index2 >= l2 || arr1[index1] < arr2[index2]) {
			ret = append(ret, arr1[index1])
			index1++
		} else {
			ret = append(ret, arr2[index2])
			index2++
		}
	}

	if (l1+l2)%2 == 0 {
		return (float32(ret[len(ret)-1]) + float32(ret[len(ret)-2])) / 2
	} else {
		return float32(ret[len(ret)-1])
	}
}
