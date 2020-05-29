package sort

func CheckSort(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
	return -1
}
