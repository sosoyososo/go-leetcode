package sort

func swap(arr *[]int, a, b int) {
	tmp := (*arr)[b]
	(*arr)[b] = (*arr)[a]
	(*arr)[a] = tmp
}

func needSort(arr *[]int) bool {
	if nil == arr || len(*arr) < 2 {
		return false
	}
	return true
}
