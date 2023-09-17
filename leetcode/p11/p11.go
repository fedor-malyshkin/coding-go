package p11

/*
https://leetcode.com/problems/container-with-most-water/
*/
func maxArea(height []int) int {
	var (
		s = 0
		e = len(height) - 1
	)
	cma := min(height[s], height[e]) * (e - s)
	cmh := max(height[s], height[e])
	for s < e {
		if height[s] < cmh {
			s++
		} else {
			e--
		}
		cmh = max(height[s], height[e])
		cma = max(min(height[s], height[e])*(e-s), cma)
	}
	return cma
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
