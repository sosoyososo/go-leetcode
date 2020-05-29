package sort

func HeapSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	for i := 0; i < len(*arr); i++ {
		heapfy(arr, 0, len(*arr)-i)
		swap(arr, 0, len(*arr)-i-1)
	}
}

func heapfy(arr *[]int, index, maxIndex int) {
	if index >= maxIndex {
		return
	}

	left := index*2 + 1
	right := left + 1

	heapfy(arr, left, maxIndex)
	heapfy(arr, right, maxIndex)

	largest := index
	if left < maxIndex && (*arr)[left] > (*arr)[largest] {
		largest = left
	}
	if right < maxIndex && (*arr)[right] > (*arr)[largest] {
		largest = right
	}
	swap(arr, largest, index)
}
