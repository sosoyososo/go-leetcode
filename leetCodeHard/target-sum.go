package leetCodeHard

func findTargetSumWays(nums []int, S int) int {
	ret := 0
	step(nums, 0, func(s int) {
		if s == S {
			ret += 1
		}
	})
	return ret
}

func step(arr []int, s int, add func(int)) {
	if len(arr) == 0 {
		add(s)
		return
	}
	v := arr[0]
	step(arr[1:], s+v, add)
	step(arr[1:], s-v, add)
}
