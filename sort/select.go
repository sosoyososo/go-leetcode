package sort

func SelectSort(arr *[]int) {
	if !needSort(arr) {
		return
	}

	/**
	 * 从前K个选出一个最大的，交换他跟第K个的位置
	 * 第K个要参与排序
	 */
	for i := len(*arr) - 1; i > 0; i-- {
		maxEleIndex := 0
		for j := 0; j <= i; j++ {
			if (*arr)[j] > (*arr)[maxEleIndex] {
				maxEleIndex = j
			}
		}
		swap(arr, maxEleIndex, i)
	}
}
