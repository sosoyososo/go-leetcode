package leetCodeHard

/*
水桶问题：
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/
func maxArea(height []int) int {
	min := func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		}
		return i2
	}
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			sum := (j - i) * min(height[i], height[j])
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
