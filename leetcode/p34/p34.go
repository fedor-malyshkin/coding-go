package p34

func searchRange(nums []int, target int) []int {

	firstOccurrence := findBound(nums, target, true)

	if firstOccurrence == -1 {
		return []int{-1, -1}
	}

	lastOccurrence := findBound(nums, target, false)

	return []int{firstOccurrence, lastOccurrence}
}

func findBound(nums []int, target int, isFirst bool) int {
	N := len(nums)
	var (
		begin = 0
		end   = N - 1
	)

	for begin <= end {
		mid := (begin + end) / 2
		switch {
		case nums[mid] == target:
			if isFirst {
				// This means we found our lower bound.
				if mid == begin || nums[mid-1] != target {
					return mid
				}

				// Search on the left side for the bound.
				end = mid - 1

			} else {

				// This means we found our upper bound.
				if mid == end || nums[mid+1] != target {
					return mid
				}

				// Search on the right side for the bound.
				begin = mid + 1
			}

		case nums[mid] > target:
			end = mid - 1
		default:
			begin = mid + 1
		}
	}
	return -1
}
