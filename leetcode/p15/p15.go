package p15

import (
	"sort"
)

/*
https://leetcode.com/problems/3sum/
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result = [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			result = append(result, twoSum(nums, i)...)
		}
	}
	return result
}

func twoSum(nums []int, i int) [][]int {
	var result = [][]int{}
	var (
		lo = i + 1
		hi = len(nums) - 1
	)
	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		switch {
		case sum < 0:
			lo += 1
		case sum > 0:
			hi -= 1
		default:
			result = append(result, []int{nums[i], nums[lo], nums[hi]})
			lo += 1
			hi -= 1
			for lo < hi && nums[lo] == nums[lo-1] {
				lo += 1
			}
		}
	}
	return result
}
