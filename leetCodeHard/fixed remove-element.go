package leetCodeHard

func removeElement(nums []int, val int) int {
	validLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[validLen] = nums[i]
			validLen++
			continue
		}
	}
	return validLen
}
