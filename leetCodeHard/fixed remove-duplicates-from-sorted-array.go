package leetCodeHard

func removeDuplicates(nums []int) int {
	s := 0
	for i, _ := range nums {
		if i == 0 {
			s++
			continue
		}
		if nums[i] == nums[i-1] {
		} else {
			nums[s] = nums[i]
			s++
		}
	}
	for i := s; i < len(nums); i++ {
		nums[i] = 0
	}

	return s
}
