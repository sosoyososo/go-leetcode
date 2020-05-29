package leetCodeHard

/*
 * 1. 从后向前判断是否是最大，如果是最大，就向前一位，直到整个序列最大，或者非最大；
 * 2. 非最大数列的一个特征是最大排序不是从小到大，一定存在某个元素小于后面元素；
 * 3. 若整个数列为最大数列，就对数列进行翻转
 * 4. 如果非最大数列，将之变为相对目前次大数列。大元素提前即可办到；
 */
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	right := len(nums) - 1
	leftIndex := len(nums) - 2
	for i := 0; i < len(nums); i++ {
		left := leftIndex - i
		if isMaxPermutation(&nums, left, right) {
			continue
		} else {
			makeBigger(&nums, left, right)
			return
		}
	}
	makeBigger(&nums, 0, right)
}

func isMaxPermutation(nums *[]int, left, right int) bool {
	if nums == nil {
		return true
	}
	if len(*nums) <= 1 {
		return true
	}
	if left >= right {
		return true
	}
	for i := left; i < right; i++ {
		if (*nums)[left] >= (*nums)[right] {
			continue
		}
		return false
	}
	return true
}

/**
 * 1. 最大的话，翻转
 * 2. 非最大，后面的是最大，需要找到右边刚好大于left位置的数字，然后剩余数字排序
 */
func makeBigger(nums *[]int, left, right int) {
	if (*nums)[left] > (*nums)[left+1] {
		turnOver(nums, left, right)
		return
	}

	ele := (*nums)[left]
	turnOver(nums, left+1, right)
	index := left + 1
	minBiggerThaeEle := ele
	for i := left + 1; i <= right; i++ {
		item := (*nums)[i]
		if item > ele && item < minBiggerThaeEle {
			index = i
			break
		}
	}
	(*nums)[left] = (*nums)[index]
	(*nums)[index] = ele

	for i := left + 1; i < right; i++ {
		if (*nums)[i] > (*nums)[i+1] {
			tmp := (*nums)[i]
			(*nums)[i] = (*nums)[i+1]
			(*nums)[i+1] = tmp
		}
	}
}

func turnOver(nums *[]int, left, right int) {
	if left >= right {
		return
	}
	l := right - left + 1
	for i := 0; i < l/2; i++ {
		tmp := (*nums)[left+i]
		(*nums)[left+i] = (*nums)[right-i]
		(*nums)[right-i] = tmp
	}
}
