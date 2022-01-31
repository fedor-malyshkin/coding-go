package p713

/*
https://leetcode.com/problems/subarray-product-less-than-k/
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	ans := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		prod = prod * nums[right]
		for prod >= k {
			prod = prod / nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
