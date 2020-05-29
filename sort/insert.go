package sort

func InsertSort(arr *[]int) {
	if !needSort(arr) {
		return
	}

	for i := 1; i < len(*arr); i++ { //插入第i个元素
		for j := 0; j < i; j++ { //跟第几个元素进行对比
			if (*arr)[j] > (*arr)[i] {
				swap(arr, i, j)
			}
		}
	}
}
