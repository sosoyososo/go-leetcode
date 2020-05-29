package sort

func BubbleSort(arr *[]int) {
	if !needSort(arr) {
		return
	}

	/**
	 * 冒泡排序每一轮把最大的冒泡到最右
	 * 外层遍历的锚点可以是最大元素的位置
	 * 内层遍历的锚点可以是最后可能被交换元素的位置
	 */
	for i := 1; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}
